package tools

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

// File interface
type FileI interface {
	Path() string
}

// File struct
type File struct {
	Root      string
	Directory string
	File      string
}

// NewFile give a new/empty instance
func NewFile(root, dir, file string) *File {
	return &File{
		Root:      root,
		Directory: dir,
		File:      file,
	}
}

// Path returns the complete file path with the file suffix
func (f *File) Path() string {
	return f.Root + "/" + f.Directory + "/" + f.File + Suffix
}

// Create file
func (f *File) Create() error {
	file, err := os.Create(f.Path())
	defer file.Close()
	if err != nil {
		log.Errorf("os.Create error: err=%s", err)
		return err
	}
	return nil
}

// Sync file
func (f *File) Sync(content string) error {
	err := ioutil.WriteFile(f.Path(), []byte(content), 0644)
	if err != nil {
		log.Errorf("ioutil.WriteFile error: err=%s", err)
		return err
	}
	return nil
}

// Delete file
func (f *File) Delete() error {
	err := os.Remove(f.Path())
	if err != nil {
		log.Errorf("os.Remove error: err=%s", err)
		return err
	}
	return err
}

// For charts
type ChartsFile struct {
	File
}

func NewChartsFile(root, dir, file string) *ChartsFile {
	return &ChartsFile{
		File: *NewFile(root, dir, file),
	}
}

func (cf *ChartsFile) Path() string {
	return cf.Root + "/" + cf.Directory + "/" + cf.File.File + ChartsSuffix
}

// Mkdir make dir
func (cf *ChartsFile) Mkdir() error {
	if _, err := os.Stat(cf.Path()); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(cf.Path()), 0755)
		if err != nil {
			log.Errorf("os.Mkdir error: err=%s", err)
			return err
		}
	}
	return nil
}

// Render func
func (cf *ChartsFile) Render(obj interface{}, tmpl string) error {
	err := cf.Mkdir()
	if err != nil {
		log.Fatalf("f.Mkdir() error: path=%s, err=%s", cf.Path(), err)
	}
	file, err := os.Create(cf.Path())
	defer file.Close()
	if err != nil {
		log.Fatalf("os.Create file error: file=%s, err=%s", cf.Path(), err)
	}

	writer := bufio.NewWriter(file)
	temp := template.Must(template.New("charts").Parse(tmpl))
	err = temp.Execute(writer, obj)
	if err != nil {
		log.Errorf("template Execute error: err=%s\n", err)
		return err
	}

	writer.Flush()
	return err
}

// TemplatesFile
type TemplatesFile struct {
	File
}

func NewTemplatesFile(root, dir, file string) *TemplatesFile {
	return &TemplatesFile{
		File: File{
			Root:      root,
			Directory: dir,
			File:      file,
		},
	}
}

func (tf *TemplatesFile) Path() string {
	return tf.Root + "/" + tf.Directory + "/" + tf.File.File + "/templates"
}

func (tf *TemplatesFile) Mkdir() error {
	if _, err := os.Stat(tf.Path()); os.IsNotExist(err) {
		err = os.MkdirAll(tf.Path(), 0755)
		if err != nil {
			log.Errorf("os.Mkdir error: err=%s", err)
			return err
		}
	}
	return nil
}

// Sync func generates templates/deployment.yaml, templates/service.yaml, templates/ingress.yaml
func (tf *TemplatesFile) Sync() error {
	err := tf.Mkdir()
	if err != nil {
		log.Errorf("tf.Mkdir error: err=%s", err)
		return err
	}
	// templates/deployment.yaml
	err = ioutil.WriteFile(tf.Path()+DeploymentSuffix, []byte(DeploymentTempl), 0644)

	// templates/service.yaml
	err = ioutil.WriteFile(tf.Path()+ServiceSuffix, []byte(ServiceTempl), 0644)

	// templates/ingress.yaml
	err = ioutil.WriteFile(tf.Path()+IngressSuffix, []byte(IngressTempl), 0644)

	// templates/NOTES.txt
	err = ioutil.WriteFile(tf.Path()+NotesSuffix, []byte(NotesTxt), 0644)

	// templates/_helper.tpl
	err = ioutil.WriteFile(tf.Path()+HelperSuffix, []byte(HelperTempl), 0644)

	if err != nil {
		log.Errorf("ioutils.WriteFile deployment.yaml/service.yaml/ingress.yaml error: err=%s", err)
		return err
	}

	return nil
}

// ValuesFile
type ValuesFile struct {
	File
}

// NewValuesFile func
func NewValuesFile(root, dir, file string) *ValuesFile {
	return &ValuesFile{
		File: File{
			Root:      root,
			Directory: dir,
			File:      file,
		},
	}
}

func (vf *ValuesFile) Path() string {
	return vf.Root + "/" + vf.Directory + "/" + vf.File.File + ValuesSuffix
}

func (vf *ValuesFile) Mkdir() error {
	if _, err := os.Stat(vf.Path()); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(vf.Path()), 0755)
		if err != nil {
			log.Errorf("os.Mkdir error: err=%s", err)
			return err
		}
	}
	return nil
}

// Sync func generates templates/deployment.yaml, templates/service.yaml, templates/ingress.yaml
func (vf *ValuesFile) Render(obj interface{}, tmpl string) error {
	err := vf.Mkdir()
	if err != nil {
		log.Fatalf("f.Mkdir() error: path=%s, err=%s", vf.Path(), err)
	}
	file, err := os.Create(vf.Path())
	defer file.Close()
	if err != nil {
		log.Fatalf("os.Create file error: file=%s, err=%s", vf.Path(), err)
	}

	writer := bufio.NewWriter(file)
	temp := template.Must(template.New("values").Parse(tmpl))
	err = temp.Execute(writer, obj)
	if err != nil {
		log.Errorf("template Execute error: err=%s\n", err)
		return err
	}

	writer.Flush()
	return err
}
