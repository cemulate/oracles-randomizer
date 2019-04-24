package rom

import (
	"fmt"
	"sort"
	"strings"
)

// this file is for mutables that go at the end of banks. each should be a
// self-contained unit (i.e. don't jr to anywhere outside the byte string) so
// that they can be appended automatically with respect to their size.

// return e.g. "\x2d\x79" for 0x792d
func addrString(addr uint16) string {
	return string([]byte{byte(addr), byte(addr >> 8)})
}

// return e.g. 0x792d for "\x2d\x79"
func stringAddr(addr string) uint16 {
	return uint16([]byte(addr)[0]) + uint16([]byte(addr)[1])<<8
}

// adds code at the given address, returning the length of the byte string.
func addCode(name string, bank byte, offset uint16, code string) uint16 {
	codeMutables[name] = MutableString(Addr{bank, offset},
		string([]byte{bank}), code)
	return uint16(len(code))
}

type romBanks struct {
	endOfBank []uint16
	assembler *assembler
}

var codeMutables = map[string]Mutable{}

// appendToBank appends the given data to the end of the given bank, associates
// it with the given name, and returns the address of the data as a string such
// as "\xc8\x3e" for 0x3ec8. it panics if the end of the bank is zero or if the
// data would overflow the bank.
func (r *romBanks) appendToBank(bank byte, name, data string) string {
	eob := r.endOfBank[bank]

	if eob == 0 {
		panic(fmt.Sprintf("end of bank %02x undefined for %s", bank, name))
	}

	if eob+uint16(len(data)) > 0x8000 {
		panic(fmt.Sprintf("not enough space for %s in bank %02x", name, bank))
	}

	codeMutables[name] = MutableString(Addr{bank, eob}, "", data)
	r.endOfBank[bank] += uint16(len(data))

	return addrString(eob)
}

// appendASM acts as appendToBank, but by compiling a block of asm. additional
// arguments are formatted into `asm` by fmt.Sprintf. the returned address is
// also given as a uint16 rather than a big-endian word in string form.
func (r *romBanks) appendASM(bank byte, name, asm string,
	a ...interface{}) uint16 {
	var err error
	asm, err = r.assembler.compileBlock(fmt.Sprintf(asm, a...), ";\n")
	if err != nil {
		panic(err)
	}
	return stringAddr(r.appendToBank(bank, name, asm))
}

// replace replaces the old data at the given address with the new data, and
// associates the change with the given name. actual replacement will fail at
// runtime if the old data does not match the original data in the ROM.
func (r *romBanks) replace(bank byte, offset uint16, name, old, new string) {
	codeMutables[name] = MutableString(Addr{bank, offset}, old, new)
}

// replaceASM acts as replace, but by compiling a block of asm. additional
// arguments are formatted into `asm` by fmt.Sprintf.
func (r *romBanks) replaceASM(bank byte, offset uint16, name, old, asm string,
	a ...interface{}) {
	var err error
	asm, err = r.assembler.compileBlock(fmt.Sprintf(asm, a...), ";\n")
	if err != nil {
		panic(err)
	}
	r.replace(bank, offset, name, old, asm)
}

// replaceMultiple acts as replace, but operates on multiple addresses.
func (r *romBanks) replaceMultiple(addrs []Addr, name, old, new string) {
	codeMutables[name] = MutableStrings(addrs, old, new)
}

// returns an ordered slice of keys for slot names, so that dentical seeds
// produce identical checksums.
func getOrderedSlotKeys() []string {
	keys := make([]string, 0, len(ItemSlots))
	for k := range ItemSlots {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// returns a byte table of (group, room, collect mode) entries for randomized
// items. in ages, a mode >7f means to use &7f as an index to a jump table for
// special cases.
func makeCollectModeTable() string {
	b := new(strings.Builder)

	for _, key := range getOrderedSlotKeys() {
		slot := ItemSlots[key]

		// trees and slots where it doesn't matter (shops, rod)
		if slot.collectMode == 0 {
			continue
		}

		var err error
		if slot.collectMode == collectFall && slot.Treasure != nil &&
			slot.Treasure.id == 0x30 {
			// use falling key mode (no fanfare) if falling item is a key
			_, err = b.Write([]byte{slot.group, slot.room, collectKeyFall})
		} else {
			_, err = b.Write([]byte{slot.group, slot.room, slot.collectMode})
		}
		if err != nil {
			panic(err)
		}
	}

	b.Write([]byte{0xff})
	return b.String()
}

// returns a byte table (group, room, ID, subID) entries for randomized small
// key drops (and other falling items, but those entries won't be used).
func makeKeyDropTable() string {
	b := new(strings.Builder)

	for _, key := range getOrderedSlotKeys() {
		slot := ItemSlots[key]

		if slot.collectMode != collectFall {
			continue
		}

		// accommodate nil treasures when creating the dummy table before
		// treasures have actually been assigned.
		var err error
		if slot.Treasure == nil {
			_, err = b.Write([]byte{slot.group, slot.room, 0x00, 0x00})
		} else if slot.Treasure.id == 0x30 {
			// make small keys the normal falling variety, with no text box.
			_, err = b.Write([]byte{slot.group, slot.room, 0x30, 0x01})
		} else {
			_, err = b.Write([]byte{slot.group, slot.room,
				slot.Treasure.id, slot.Treasure.subID})
		}
		if err != nil {
			panic(err)
		}
	}

	b.Write([]byte{0xff})
	return b.String()
}
