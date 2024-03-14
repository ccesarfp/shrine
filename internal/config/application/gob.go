package application

import (
	"encoding/gob"
	status "github.com/ccesarfp/shrine/internal/enum/status"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Gob - used to store Shrine Status
type Gob struct {
	Name      string
	Version   string
	Pid       int
	Address   string
	StartTime time.Time
	Status    uint8
}

const tmpPath = "/tmp"
const folder = "shrine*"
const file = "s"

var g Gob

// write - Write Gob in temporary files
func write(app *Application) error {
	g = Gob{
		Name:      app.Name,
		Version:   app.Version,
		Pid:       os.Getpid(),
		Address:   app.s.Address,
		StartTime: app.s.StartTime,
		Status:    status.Close,
	}

	dir, err := createOrFindDir()
	if err != nil {
		return err
	}

	f, err := createOrFindGob(filepath.Join(dir, file))
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	enc := gob.NewEncoder(f)
	err = enc.Encode(g)
	if err != nil {
		return err
	}

	return nil
}

// createOrFindDir - try to find temporary directories, if fails, create the directory
func createOrFindDir() (string, error) {
	matches, err := filepath.Glob(filepath.Join(tmpPath, folder))
	if err != nil {
		return "", err
	}

	if len(matches) > 0 {
		return matches[0], nil
	}

	dir, err := os.MkdirTemp("/tmp", folder)
	if err != nil {
		return "", err
	}

	return dir, nil
}

// createOrFindGob - try to find temporary file, if fails, create the file
func createOrFindGob(file string) (*os.File, error) {
	matches, err := filepath.Glob(file + strconv.Itoa(g.Pid) + "*")
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		f, err := os.OpenFile(matches[0], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}
		return f, nil
	}

	f, err := os.Create(file + strconv.Itoa(g.Pid) + ".gob")
	if err != nil {
		return nil, err
	}

	return f, nil
}

// findGobs - find the name of multiples gobs
func findGobs() ([]string, error) {
	matches, err := filepath.Glob(filepath.Join(tmpPath, folder, file+"*"))
	if err != nil {
		return nil, err
	}

	return matches, nil
}

// openGob - open Gob
func openGob(file string) (*os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Read - read data inside of multiple Gobs
func Read() ([]Gob, error) {
	gobsNames, err := findGobs()
	if err != nil {
		return nil, err
	}

	var gobs []Gob
	for _, gobName := range gobsNames {
		var g Gob
		f, err := openGob(gobName)
		if err != nil {
			return nil, err
		}
		dec := gob.NewDecoder(f)
		err = dec.Decode(&g)
		if err != nil {
			return nil, err
		}
		_ = f.Close()
		gobs = append(gobs, g)
	}

	return gobs, nil
}

// remove - remove gob
func remove() error {
	files, err := filepath.Glob(filepath.Join(tmpPath, folder, file+strconv.Itoa(g.Pid)+"*"))
	if err != nil {
		return nil
	}

	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}

	return nil
}
