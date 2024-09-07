# NASTeaPi

**__THIS IS STILL A VERY WIP BUILD. I WOULD WAIT UNTIL THIS IS REMOVED BEFORE 3D PRINTING OR BUILDING THIS__**



![image](https://github.com/user-attachments/assets/d0cc242d-6f9c-4613-ae51-bb5666934720)

NASTea Pi, a small pi powered NAS that is cheap to build, but as the saying goes, it's not exactly fast to build, and it's not overly that great either (see notes). 

Prerequisite Tools:

- Soldering Iron (and knowledge to use it)
- Screwdriver of your choice (Pair with the M3/M4 screws you use later)
- 3D Printer

Part List:

- 1x Sparkfun USB-C Breakout
- 3x 2.5inch 5400RPM SATA Hard Drives (Or SATA SSDs)
- 1x Raspberry Pi (2,3,4, and 5 are supported but I'll be using a 3 for the guide)
- 3x SATA to USB Adapters
- 5x M4x6 Screws (x8 will work too)
- 4x M3x30 Screws
- 4x M3 Nuts
- 1x 40x40x20mm 5V PWM Fan
- 2x 5.1k ohm resistors
- 24ga wire for data
- 16ga wire for power

Notes:
I want this as a secondary site NAS that just sits in a corner and happily backs up data from my Synology disk station. It's not designed as a primary network storage device but I do want to come back around for a second revision and build a better one with better support for using drives at their full speed. This is wired in a way that you are limited to USB 2.0 Link speeds (even if you buy USB3.0 adapters). The PI4 and 5 do not have enough USB3 ports to connect all 3 drives, and the 2 and 3 only have 2.0 ports, so if another revision is done it will have to include a hub to act as a controller and be restricted to the 4 and 5 model. Being that this has already gotten a bit out of hand scope wise, I didn't really feel like adding a custom PCB to the mix would be helpful. 

Also, for the adapters, I used [these ones](https://www.amazon.com/dp/B073SXTY64) and then you carefully crack the casing open to get the PCB. There are a few similar designs on Amazon and they all seem to differ slightly on the PCB design so just keep that in mind. I tried to design it so that it would fit any adapter PCBs that have a similar layout of link/activity LEDs with the small casing. Your mileage may vary but try as I might I couldn't find anyone who just sold these PCBs that I could buy from. 

Also also, I recommend a [Noctua 40x40x20mm](https://www.amazon.com/gp/product/B071FNHVXN) fan (they have a 5v 4-pin version that I used) as they are MUCH quieter than the alternatives I tried. As for the USB C Breakout, any with [this rough layout](https://www.amazon.com/dp/B07M6R37L8) will do just fine. 

![image](https://github.com/user-attachments/assets/1edec0ec-a0c8-4e96-a2f0-7c502580d46a)

Instructions and wiring diagrams to be done in the coming days after I finish the initial build. 
