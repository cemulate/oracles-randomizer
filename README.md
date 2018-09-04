# Oracle of Seasons randomizer

This program reads a Zelda: Oracle of Seasons ROM (US version only) shuffles
the locations of items and mystical seeds, randomizes the default season for
each area, and writes the modified ROM to a new file. It also bypasses essence
checks for overworld events that are necessary for progress, so the dungeons
can be done in any order that the randomized items facilitate. However, you do
have to collect all 8 essences to get the Maku Seed and finish the game.


## Usage

There are three ways to use the randomizer:

1. Place the randomizer in the same directory as your vanila OoS ROM (or vice
   versa), and run it. The randomizer will automatically choose the vanilla ROM
   and write the randomized ROM and log to a new file.
2. In Windows, drag your ROM onto the executable. If the ROM is vanilla, it
   will be randomized; otherwise it will be updated to the latest version,
   applying bugfixes and other changes when applicable. In either case the
   original file is not overwritten.
3. Use the command line. Type `oos-randomizer -h` to view the usage summary.
   The `-freewarp` and `-seed` options are probably the only useful end-user
   ones.


## Download

You can download executables for Windows, macOS, and Linux from the
[releases](https://github.com/jangler/oos-randomizer/releases) page.


## Randomization notes

Items and chests are randomied, with exceptions listed below. The rod of
seasons is split into four items, each of which will give you one season and
the rod itself (if you don't already have it). There is one flute in the game
for a random animal companion, and it's identified and usable as soon as you
get it. Subrosian dancing and Ricky do not give flutes as they normally would.

Seed trees and default seasons for each area are also shuffled, and the satchel
and slingshot will start with the type of seeds on the tree in Horon Village.
The Natzu region matches whichever animal companion the randomized flute calls.

The randomizer will never require you to farm rupees if you spend them wisely.

The following items are **not** randomized:

- Renewable shop items (bombs, shield, hearts, etc.)
- Small keys
- Pirate's bell
- Found items (gasha seeds and pieces of heart outside of chests)
- Subrosian dancing prizes after the first
- Trading sequence items
- Non-essential items given by NPCs
- Subrosian hide and seek items
- Gasha nut contents
- Fixed drops
- Maple drops


## Other notable changes

Other small changes have been made for convenience, to simplify randomization
logic, or to prevent softlocks. The most notable are:

- The intro sequence and pirate cutscene are removed.
- Mystical seeds grow in all seasons.
- Seeds can be collected if the player has either a slingshot or the satchel.
- The cliff between Eastern Suburbs and Sunken City is blocked except in
  spring.
- Rosa doesn't appear in the overworld, and her portal is activated by default.
- Fool's ore is randomized (the Strange Brothers trade you nothing for your
  feather).
- The diving spot at the south end of Sunken City is removed.
- **Holding start while closing the map screen outdoors warps to the seed tree
  in Horon Village.** This also sets your save/respawn point to that screen.
  Tree warping has a one-hour cooldown unless the `-freewarp` flag is
  specified. Tree warp is not supported as a "feature" and has no warranty, so
  consider possible consequences before using it.

## FAQ

**Q: Do I have to do HSS skip or Poe skip?**

A: No, but you can if you want to, and the randomizer accounts for those and other
sequence breaks. Though depending on the seed, you might be expected to do any
number of tricky things the vanilla game wouldn't expect of you.

**Q: I'm softlocked. Now what do I do?**

A: If you're softlocked by location, use tree warp. Otherwise, open an issue
about it or tell me in Discord, and provide the log file. Depending on the
problem, you may be able to `-update` your ROM using the next patch version to
un-softlock.

**Q: Are you going to make a randomizer for Oracle of Ages too?**

A: Maybe, but not until the Seasons randomizer is reasonably feature-complete
(as i see it). Ages also has some big sequence breaks and Crescent Island,
which would both be tricky to account for in the logic unless they're just
removed.

**Q: Is there a place to discuss the randomizer?**

A: Yes, the Oracles Discord server (link
[here](https://www.speedrun.com/oos/thread/3qwe1)). The server is mainly focused
on speedrunning, but the #randomizer channel is for anything pertaining to the
randomizer.
