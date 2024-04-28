package export

import (
	"bufio"
	"encoding/json"
	"os"
)

type ExportJsonMethod struct {
	Data any
}

func (ejm ExportJsonMethod) Export() error {
	data, err := json.Marshal(ejm.Data)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Seek(0, 0); err != nil {
		return err
	}

	if err = f.Truncate(0); err != nil {
		return err
	}

	writter := bufio.NewWriter(f)

	if _, err = writter.Write(data); err != nil {
		return err
	}

	if err = writter.Flush(); err != nil {
		return err
	}

	println("Exporting json success")
	return nil
}
