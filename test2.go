package main

import (
    "os"
    "path/filepath"
    "github.com/godbus/dbus"
    "github.com/linuxdeepin/go-lib/log"
    dutils "github.com/linuxdeepin/go-lib/utils"
    "github.com/linuxdeepin/go-lib/imgutil"
    "github.com/linuxdeepin/go-lib/strv"
)

//export DDE_DEBUG_MATCH = test
var logger = log.NewLogger("test")

func getLicenseAuthorizationProperty() uint32 {
    conn, err := dbus.SystemBus()
    if err != nil {
        logger.Warning(err)
        return 0
    }
    var variant dbus.Variant
    err = conn.Object("com.deepin.license", "/com/deepin/license/Info").Call(
        "org.freedesktop.DBus.Properties.Get", 0, "com.deepin.license.Info", "AuthorizationProperty").Store(&variant)
    if err != nil {
        logger.Warning(err)
        return 0
    }

    logger.Warning("111111", variant.Signature().String())
    if variant.Signature().String() != "u" {
        logger.Warning("not excepted value type")
        return 0
    }

    return variant.Value().(uint32)
}

func main() {
    logger.Warning(getLicenseAuthorizationProperty())

    logger.Warning("---- getSysBgFiles :", getSysBgFiles(0))

    logger.Warning("---- getSysBgEnterpriseFiles :", getSysBgFiles(1))

    logger.Warning("---- getSysBgGovernmentFiles :", getSysBgFiles(2))

    logger.Warning("---- getSysBgGovernmentFiles :", getSysBgFiles(5))
}

var uiSupportedFormats = strv.Strv([]string{"jpeg", "png", "bmp", "tiff", "gif"})

func IsBackgroundFile(file string) bool {
    file = dutils.DecodeURI(file)
    format, err := imgutil.SniffFormat(file)
    if err != nil {
        return false
    }

    if uiSupportedFormats.Contains(format) {
        return true
    }
    return false
}

func getBgFilesInDir(dir string) []string {
    fr, err := os.Open(dir)
    if err != nil {
        return []string{}
    }
    defer fr.Close()

    names, err := fr.Readdirnames(0)
    if err != nil {
        return []string{}
    }

    var walls []string
    for _, name := range names {
        path := filepath.Join(dir, name)
        if !IsBackgroundFile(path) {
            continue
        }
        walls = append(walls, path)
    }
    return walls
}

var (
    systemWallpapersDir = []string {
        "/usr/share/wallpapers/deepin",
        "/usr/share/wallpapers/deepin/deepin-enterprise",
        "/usr/share/wallpapers/deepin/deepin-government",
    }
)

func getSysBgFiles(licenseType int) []string {
    var files []string
    if licenseType < len(systemWallpapersDir)  {
        files = append(files, getBgFilesInDir(systemWallpapersDir[licenseType])...)
    }
    return files
}