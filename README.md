# Screenshot Renamer for macOS (Auto-Startup)

This project automatically renames macOS screenshots by removing spaces and formatting them consistently.

## **Installation & Auto-Startup Setup**
Follow these steps to install and configure the screenshot renamer to run automatically on macOS startup.

---

### Move the Compiled Binary to a Permanent Location

**To ensure the binary is in a safe and accessible location, move it to `~/Applications`:**

```bash
mkdir -p ~/Applications
mv rename_screenshots ~/Applications/
chmod +x ~/Applications/rename_screenshots
```

**Make a user launch agents directory and edit the file:**
```bash
mkdir -p ~/Library/LaunchAgents
vim ~/Library/LaunchAgents/com.screenshot.renamer.plist
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>Label</key>
        <string>com.screenshot.renamer</string>

        <key>ProgramArguments</key>
        <array>
            <string>/Users/YOUR_USERNAME/Applications/rename_screenshots</string>
        </array>

        <key>RunAtLoad</key>
        <true/>

        <key>KeepAlive</key>
        <true/>
    </dict>
</plist>
```

