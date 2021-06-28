# bifocals
Linux version of Spectacle for macOS

# Install
1) Clone the repo and run in its folder:
   `make && sudo make install`
2) Add as many of the keyboard shortcuts as you please:
   `bifocals left`
   `bifocals right`
   `bifocals top`
   `bifocals bottom`
   `bifocals fullscreen`
3) Optional - Remove the default shortcuts:
   `View split on left`
   `View split on right`
   `Maximize window`
   `Restore window`

# Issues
* Untested on more than 2 monitors and many configurations
* If you have used `cmd+left`, `cmd+right`, or other Ubuntu window resize shortcut, then you must "detach" it before you can use these shortcuts on that window.

# TODO
* Ensure dependencies are installed before allowing to `make install`
* Create a `.deb` package 