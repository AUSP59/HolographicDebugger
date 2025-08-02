
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "io"
    "os"
)

// ComputeFileSHA256 returns the SHA-256 hash of the given file.
func ComputeFileSHA256(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()

    h := sha256.New()
    if _, err := io.Copy(h, f); err != nil {
        return "", err
    }

    return hex.EncodeToString(h.Sum(nil)), nil
}
