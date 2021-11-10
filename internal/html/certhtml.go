package html

import (
	"certificat-go/internal/cert"
	"fmt"
	"html/template"
	"os"
	"path"
)

type HtmlSaver struct {
	OutPutDir string
}

func New(outPutdir string) (*HtmlSaver, error) {
	var h *HtmlSaver

	err := os.MkdirAll(outPutdir, os.ModePerm)
	if err != nil {
		return h, err
	}

	h = &HtmlSaver{
		OutPutDir: outPutdir,
	}

	return h, nil

}

func (p *HtmlSaver) Save(cert cert.Cert) error {
	t, err :=template.New("certificate").Parse(tpl)
	if err !=nil {
		return err
	}
	fileName :=fmt.Sprintf("%v.html", cert.LabelTitle)
	path :=path.Join(p.OutPutDir,fileName)
	f, err := os.Create(path)
	if err !=nil {
		return err
	}
	defer f.Close()
	err = t.Execute(f,cert)
	if err !=nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v' \n", path)
	return nil
}

var tpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.LabelTitle}}</title>
</head>
<body>
    <h1>{{.LabelCompletion}}</h1>
	<h2><em>{{.LabelPresented}}</em></h2>
	<h1>{{.Name}}</h1>
	<h2>{{.LabelParticipation}}</h2>
	<p><em>{{.LabelDate}}</em></p>

</body>
<style>
 body {
	 display:flex;
	 flex-direction: column;
	 justify-content : center;
	 text-align:center;
 }
</style>
</html>`