package utils

import (
  "crypto/rand"
  "encoding/hex"
)

func GenerateRandomHash(size int) string {
  if size == 0 {
    return "";
  }

  bytes := make([]byte, size);
  if _, err := rand.Read(bytes); err != nil {
    return "";
  }
  return hex.EncodeToString(bytes);
}
