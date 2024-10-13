package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {
	tp, err := readTemplate([]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"}, templateDir)

	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var templateBlogs []TemplateBlog
	for _, templateName := range templates {
		templateName = templateName + ".html"
		t := template.New(templateName)
		t_main := templateDir + "home.html"
		t_header := templateDir + "layout/header.html"
		t_footer := templateDir + "layout/footer.html"
		t_personal := templateDir + "layout/personal.html"
		t_post_list := templateDir + "layout/post-list.html"
		t_pagination := templateDir + "layout/pagination.html"
		t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": date, "dateDay": dateDay})
		t, err := t.ParseFiles(templateDir+templateName, t_main, t_header, t_footer, t_personal, t_post_list, t_pagination)
		if err != nil {
			log.Println("模板解析出错", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = t
		templateBlogs = append(templateBlogs, tb)
	}
	return templateBlogs, nil
}
func isODD(idx int) bool {
	return idx%2 == 0
}
func getNextName(strs []string, idx int) string {
	return strs[idx+1]
}
func date(format string) string {
	return time.Now().Format(format)
}
func dateDay(t time.Time) string {
	return t.Format("2006-01-02 15:04:06")
}
func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
