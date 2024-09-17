# NASTeaPi

NASTea Pi, a small pi powered NAS that is cheap to build, but as the saying goes, it's not exactly fast to build, and it's not overly that great either. For me the purpose of this was to act as a secondary data backup site as I have a Synology disk station for my main backups. The priority on this was low power and redundancy so that my other NAS can send backups to it overnight. I do plan on a 2.0 of this with a more robust design.
![IMG_20240916_134336444](https://github.com/user-attachments/assets/bc13ccd7-e96a-4e18-a629-dc6ed73f86f8)

# Introduction
I want to preface this whole thing and potentially save you some reading by saing: In reality, there are a lot of better projects that do the same function with better documentation, better design, and better compatibility, however this was kind of purpose built to use up what I have on-hand. I ended up making 2 of these, both of the internal designs were different but I will be detailing the one that ended up needing less work (still a lot of work). See the end of this doc for the full gallery of images. 

This chassis is made to be relatively flexible and also I've designed front plates for:
- RPi 4/5
- RPi 1/2/3
- Modified RPi 1/2/3


# Parts
**Prerequisite Tools:**
- Soldering Iron (and knowledge to use it)
- Solder
- Flux
- Hot air rework station (optional, for removing USB ports)
- Kapton Take (optional, for removing USB ports)
- Desoldering Wick (optional, for removing USB ports)
- Pliers (optional, for removing USB ports)
- Vice (optional, for removing USB ports)
- Screwdriver of your choice (Pair with the M3/M4 screws you use later)
- 3D Printer
- Shrink Tube

**Part List:**
- 1x Sparkfun USB-C Breakout
- 1x Female USB A 2.0 Through Hole Connector (Mouser: 87520-3010ALF) 
- 3x 2.5inch _LOW POWER_ SATA Hard Drives (Or SATA SSDs)
- 1x Raspberry Pi (1,2,3,4, and 5 are supported but I'll be using a 3b for the guide)
- 3x SATA to USB Adapters
- 5x M4x6 Screws (x8 will work too)
- 4x M3x30 Screws
- 4x M3 Nuts
- 1x 40x40x20mm 5V PWM Fan
- 2x 5.1k ohm resistors
- 24ga wire (for power and usb)
- 26ga wire (for USB if you prefer)

Note for the drives: 
It's important to know how much power your drives need. The drives I'm using are some Kingston SSDs I had laying around, these draw about 1.2 watts each during writes. Whatever drives you use need to draw less than 5 watts combined, otherwise it will overload the USB ports (they can do 6, but you have to account for the adapters too). If they draw MORE, I have a few things you can do to give you a bit more headroom. 

Note on the adapters:
I used [these ones](https://www.amazon.com/dp/B073SXTY64) and then you carefully crack the casing open to get the PCB. There are a few similar designs on Amazon and they all seem to differ slightly on the PCB design so just keep that in mind. I tried to design it so that it would fit any adapter PCBs that have a similar layout of link/activity LEDs with the small casing. Your mileage may vary but try as I might I couldn't find anyone who just sold these PCBs that I could buy from. 

Note on the fan:
I recommend a [Noctua 40x40x20mm](https://www.amazon.com/gp/product/B071FNHVXN) fan (they have a 5v 4-pin version that I used) as they are MUCH quieter than the alternatives I tried. 

Note on the USB-C Breakout:
I just went with a Sparkfun breakout, but any with [this rough layout](https://www.amazon.com/dp/B07M6R37L8) will do just fine. 


# Build
I'll be doing a better write-up on this in the future, might make a video or a blog post going through the steps, but here is the rough guide. Again, I am not being very detailed here and you could easily burn yourself doing this so do NOT attempt this unless you know what you're doing. 

So, first up you gotta 3D Print everything, the parts are all designed to be printed without supports in this orientation:
![image](https://github.com/user-attachments/assets/b58eb38f-0422-4726-bb30-df61f981468e)

Once you have everything printed, you may need to file some rough edges but overall very smooth. 

This part is optional, but it is what I did to make soldering the connectors easier AND to improve airflow over the CPU. So time to take your raspberry PI 3, and desolder the USB connectors (all of them). First, put kapton tape on the bottom of the PI in a way that will protect the small SMD components on the board near the USB ports, then place the PI gently into a vice. With about 420 degrees of heat, gently use a rework station to blow over the USB pins of each set of USB ports while putting light pulling pressure with some pliers. After a few moments they will come free, you may have some internal pins get stuck, just desolder those with a normal iron. Once you have them out, use some solder wick to pull out the rest of the solder. 

Now, if your drives use LESS than 5w total power, then you can solder wires directly to the solder points for the USB ports following this pinout.
![image](https://github.com/user-attachments/assets/a7d29cb5-6e4e-426d-8b99-8c05b4f17c50)

If your drives use MORE than 5w total power, you can try tapping the V+ straight from the USB-C adapter (This is how the PiNAS project does it), however since most adapters only deliver 3A max, this may not work and you're likely better off finding another project that integrates power better. For my USB 4 powered one, I ended up stuffing a powered USB adapter in the case too and fed another cable through the front to get power to it. Jank, but it works. 

Once you have everything soldered up to the PI, REMOVE the USB cables from the Sata adapters you have (MAKE SURE TO TAKE A PICTURE OF THE PINOUT) then solder the cables to it. I made 5 inch cables that were 24awg stranded wire so that they would be less delicate, I could have used thinner wire though. 

Next, solder the USB-C breakout to the two test pads near your Raspberry PI's power input. (They are all different so this is a, you should google your specific pinout situation). It's optional, but I'd also solder 5.1k ohm resistors from CC1 and CC2 to ground, this tells PD chargers to send 5v power, but since most will default to this anyway, you can probably skip that step. I did on the second unit and it's been fine.

Lastly, take your female USB connector, and solder it to the place I marked above as "DO NOT USE THIS PORT"

After that, you should have something that looks like this:
![IMG_20240912_143647902](https://github.com/user-attachments/assets/bd32660b-a6bb-4371-92dd-a7fc859ce0fd)

We'll continue the build in the "Assembly" below after software.

# Software

So honestly, I'm still figuring out how I want to do this, but so far I'm just doing everything manually with mdadm to create the raid and the normal suite of services to create the shares and whatnot. I tried getting this to work in open media vault but the raid options are lackluster since it doesn't seem to support RAID 5 out of the box. 

I will enhance this as I deploy these and use them. For now, you'll want to look into options that work for you and test in the meantime. Flash Raspberry Pi OS Lite (64Bit) to the pi, then once it's setup, use the `lsblk` command. You should see `sda, sdb, and sdc` like so: 
![image](https://github.com/user-attachments/assets/1f79fc0f-e1b0-4dae-bbb0-12b0e283f7ff)

If you see this, then at least you know your soldering job is done, if you DONT see these with drives connected then your USB connection needs another look. 

I also have a fan speed service included in here but I need to test/document it more.


# Assembly
Basically, we're just shoving everything in. Carefully put the SATA adapters in through the back while sliding drives in the front. Line them up and push them in place. Then we HOT GLUE the Raspberry PI in place (I know, jank, just make sure to put the front panel on to help get it aligned). Once that's dry, we're going to glue the USB connector to the side (I used a q-tip to hold it in place while it cured).

After that, put the M3x30 Screws through the fan grille, fan, and rear shroud, using a nut on the other end to hold them all together, this should make a little fan sandwich. Then we take the M4 screws and screw the rear shroud onto the chassis. In the meantime before I get the documentation out, you can solder or make a connector for the V+ and gnd on the fan to go to the GPIO, this way it's working and active (though always at 100 percent)
![IMG_20240916_154040939](https://github.com/user-attachments/assets/c80c4a70-3a3a-463e-9947-3e664e08d2da)

Finally slide the front panel into place, and screw it to the front of the chassis. With that, you've got yourself a NAS that isn't super fast but gets the job done! Here are all the pictures of the RPi 3 build:

![IMG_20240916_134336444](https://github.com/user-attachments/assets/dca5c1db-3f0b-47af-b2c7-ae82019df462)
![IMG_20240916_154040939](https://github.com/user-attachments/assets/ddf4cc68-791b-41c8-88f8-1a099ee67de3)
![IMG_20240916_154048385](https://github.com/user-attachments/assets/e5d0fce2-bf6a-4bfb-b2e7-af466d7c62a8)
![IMG_20240916_154416569](https://github.com/user-attachments/assets/167c3795-3314-4d82-a1cd-456b383f2540)
![IMG_20240916_154432318](https://github.com/user-attachments/assets/659543bf-790c-4305-80c3-f09e8c60eb43)


# PI 4 Build
Here are images of the other build just to get a better sense of size/scale/design in your space.
![20240906_113234](https://github.com/user-attachments/assets/18bd2b58-22e7-464c-9a11-35405c446335)
![20240906_113246](https://github.com/user-attachments/assets/021f5fb2-a567-4727-993f-0b9f75b50976)
![20240906_113255](https://github.com/user-attachments/assets/a314fa8b-b3e6-4227-819d-9241c52b5ce1)
![20240906_113307](https://github.com/user-attachments/assets/dbbc4bf9-869a-466e-9687-599e92056e89)
![20240906_113350](https://github.com/user-attachments/assets/0e2cbaa7-e6c3-477c-ac85-01bba730f467)
![20240906_113401](https://github.com/user-attachments/assets/3de29452-d0ca-4421-b9fb-bfd538f9642b)

# Designs
![image](https://github.com/user-attachments/assets/d0cc242d-6f9c-4613-ae51-bb5666934720)
![image](https://github.com/user-attachments/assets/1edec0ec-a0c8-4e96-a2f0-7c502580d46a)
