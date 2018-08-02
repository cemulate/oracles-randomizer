# Oracle of Seasons randomizer

This program reads a Zelda: Oracle of Seasons ROM (probably JP version only; I
haven't tested on EN/EU), shuffles the locations of key items and seeds, and
writes the modified ROM to a new file. It also bypasses essence checks for
overworld events that are necessary for progress, so the dungeons can be done
in any order that the randomized items facilitate. However, you do have to
collect all 8 essences to get the Maku Seed and finish the game.

The randomizer is relatively new, so consider it "beta" for now. See the
[issue tracker](https://github.com/jangler/oos-randomizer/issues) for known
problems.


## Usage

The randomizer uses a command-line interface, and I currently have no plans to
implement a graphical one. It's a simple program (from the user's perspective),
and command lines are not very hard.

The normal usage is `./oos-randomizer oos_original.gbc oos_randomized.gbc` (or
whatever filenames you want), but there are additional flags you can pass
before the filename arguments to fine-tune the randomization, as displayed in
the usage (`./oos-randomizer -h`) message:

    Usage of ./oos-randomizer:
      -devcmd string
        	if given, run developer command
      -dryrun
        	don't write an output file for any operation
      -forbid string
        	comma-separated list of nodes that must not be reachable
      -goal string
        	comma-separated list of nodes that must be reachable (default "done")
      -maxlen int
        	if >= 0, maximum number of slotted items in the route (default -1)

Note that some combinations of these flags can result in impossible conditions,
like `-goal 'd1 essence' -forbid 'ember seeds'`. See further below for an
abbreviated list of possible `-goal` and `-forbid` nodes.

Regardless of the value of `-maxlen`, the randomizer will place items in all
available slots. The flag just limits the number of slotted items that are
*necessary* in order to reach the goal(s).


## Download

You can download executables for Windows, MacOS, and Linux from the
[releases](https://github.com/jangler/oos-randomizer/releases) page.


## Randomization notes

Most inventory items (equippable and non-equippable) that are necessary to
complete a casual playthrough are shuffled, with some exceptions:

- Purchasable items (bombs, shield, and strange flute) are not shuffled.
- The ribbon and pirate's bell are not shuffled.

**Items are only placed in locations where you would normally obtain another
key item.**

Seed trees are also shuffled, and the satchel and slingshot will start with the
type of seeds on the tree in Horon Village. Remember that each type of seed
(except for mystery seeds) only grows in one season, which may not always be
the default season for the region the tree is in.

Speedrunners should note that the Subrosian dancing prize could be important.


## Potentially useful goal/forbid nodes

- `done`, meaning defeating Onox
- `dX essence`, where X is a number from 1-8
- `sword L-1`
- really just look at the files in the "prenode" folder if you want more

Remember that you can specify multiple nodes as goals/forbids by separating
them with commas. Also remember to quote the strings lol
