package main

import (
    "github.com/gin-gonic/gin"
    "github.com/go-echarts/go-echarts/v2/charts"
    "github.com/go-echarts/go-echarts/v2/opts"
    "math/rand"
    "net/http"
)

// generate random data for bar chart
func generateBarItems() []opts.BarData {
    items := make([]opts.BarData, 0)
    for i := 0; i < 7; i++ {
        items = append(items, opts.BarData{Value: rand.Intn(300)})
    }
    return items
}

func main() {
    // create a new bar instance
    bar := charts.NewBar()
    // set some global options like Title/Legend/ToolTip or anything else
    bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
        Title:    "My first bar chart generated by go-echarts",
        Subtitle: "It's extremely easy to use, right?",
    }))

    // Put data into instance
    bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
        AddSeries("Category A", generateBarItems()).
        AddSeries("Category B", generateBarItems())

    barSnippet := renderToHtml(bar)
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Home",
            "chart": barSnippet,
        })
    })
    //
    // r.GET("/graphs", func(c *gin.Context) {
    //     c.HTML(http.StatusOK, "index.")
    // })
    r.Run()
}
