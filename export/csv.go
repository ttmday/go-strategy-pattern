package export

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/yukithm/json2csv"
)

type ExportCsvMethod struct {
	Data any
}

func (ecm ExportCsvMethod) Export() error {
	var obj any
	if b, err := json.Marshal(ecm.Data); err != nil {
		return err
	} else {
		err := json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
	}

	csv, err := json2csv.JSON2CSV(obj)
	if err != nil {
		return err
	}

	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)
	file, err := os.OpenFile("data.csv", os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.Seek(0, 0); err != nil {
		return err
	}

	if err = file.Truncate(0); err != nil {
		return err
	}

	err = wr.WriteCSV(csv)
	if err != nil {
		return err
	}

	wr.Flush()
	got := b.String()

	if _, err = file.WriteString(got); err != nil {
		return err
	}

	println("Exporting csv success")
	return nil
}
