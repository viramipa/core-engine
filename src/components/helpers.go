package core_engine

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type RandomGenerator struct{}

func (r *RandomGenerator) GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	if _, err := rand.Read(uuid); err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return hex.EncodeToString(uuid), nil
}

type FilesystemHelper struct{}

func (f *FilesystemHelper) IsFileExisting(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

func (f *FilesystemHelper) IsDirectoryExisting(directoryPath string) bool {
	if _, err := os.Stat(directoryPath); err != nil {
		return false
	}
	return true
}

func (f *FilesystemHelper) CreateDirectory(directoryPath string) error {
	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (f *FilesystemHelper) RemoveDirectory(directoryPath string) error {
	if err := os.RemoveAll(directoryPath); err != nil {
		return err
	}
	return nil
}

func (f *FilesystemHelper) GetFileSize(filePath string) (int64, error) {
	if file, err := os.Stat(filePath); err != nil {
		return 0, err
	} else {
		return file.Size(), nil
	}
}

func (f *FilesystemHelper) ReadFile(filePath string) ([]byte, error) {
	if file, err := os.Open(filePath); err != nil {
		return nil, err
	} else {
		defer file.Close()
		var buffer [4096]byte
		var bytesRead int
		var data []byte
		for {
			if bytesRead, err = file.Read(buffer[:]); err != nil {
				break
			}
			data = append(data, buffer[:bytesRead]...)
		}
		return data, nil
	}
}

func (f *FilesystemHelper) WriteFile(filePath string, data []byte) error {
	if file, err := os.Create(filePath); err != nil {
		return err
	} else {
		defer file.Close()
		if _, err := file.Write(data); err != nil {
			return err
		}
		return nil
	}
}

type NetworkHelper struct{}

func (n *NetworkHelper) GetLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}

func (n *NetworkHelper) GetHostname() (string, error) {
	return os.Hostname()
}

type Logger struct{}

func (l *Logger) LogInfo(message string) {
	log.Println(fmt.Sprintf("INFO: %s", message))
}

func (l *Logger) LogError(message string) {
	log.Println(fmt.Sprintf("ERROR: %s", message))
}

type TimeHelper struct{}

func (t *TimeHelper) GetCurrentTime() time.Time {
	return time.Now()
}

func (t *TimeHelper) GetTimeInUTC() time.Time {
	return time.Now().UTC()
}

func (t *TimeHelper) GetTimeInMilliseconds() int64 {
	return t.GetCurrentTime().UnixNano() / 1e6
}

func (t *TimeHelper) FormatTime(time time.Time, layout string) string {
	return time.Format(layout)
}

type StringHelper struct{}

func (s *StringHelper) IsEmptyString(target string) bool {
	return target == ""
}

func (s *StringHelper) IsNonEmptyString(target string) bool {
	return !s.IsEmptyString(target)
}

func (s *StringHelper) IsEqualString(target1, target2 string) bool {
	return target1 == target2
}

func (s *StringHelper) IsNotEqualString(target1, target2 string) bool {
	return !s.IsEqualString(target1, target2)
}

func (s *StringHelper) IsStartsWithString(target string, prefix string) bool {
	return strings.HasPrefix(target, prefix)
}

func (s *StringHelper) IsEndsWithString(target string, suffix string) bool {
	return strings.HasSuffix(target, suffix)
}

func (s *StringHelper) GetFileNameFromPath(path string) string {
	return filepath.Base(path)
}