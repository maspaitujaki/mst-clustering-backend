package visualisasi

import (
	"backend/mst"
	"log"
	"math"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/moreland"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func MakeScatter(xAtr string, yAtr string, clusters [][]mst.Node, filePath string) {

	// for j, cluster := range clusters {
	// 	for _, Node := range cluster {
	// 		fmt.Printf("%s X:%f Y:%f CLUSTER:%d\n", Node.Name, Node.X, Node.Y, j)
	// 	}
	// }

	// randomTriples returns some random but correlated x, y, z triples
	mapTriples := func(clusters [][]mst.Node) plotter.XYZs {
		data := make(plotter.XYZs, 0)
		for i, cluster := range clusters {
			for _, node := range cluster {
				data = append(data, plotter.XYZ{X: float64(node.X), Y: float64(node.Y), Z: float64(i + 1)})
			}
		}
		return data
	}

	scatterData := mapTriples(clusters)

	// Calculate the range of Z values.
	minZ, maxZ := math.Inf(1), math.Inf(-1)
	for _, xyz := range scatterData {
		if xyz.Z > maxZ {
			maxZ = xyz.Z
		}
		if xyz.Z < minZ {
			minZ = xyz.Z
		}
	}

	// for _, data := range scatterData {
	// 	fmt.Println(data.X, data.Y, data.Z)
	// }

	colors := moreland.Kindlmann() // Initialize a color map.
	colors.SetMax(maxZ + 1)
	colors.SetMin(minZ)

	p := plot.New()
	p.Title.Text = "Cluster"
	p.X.Label.Text = xAtr
	p.Y.Label.Text = yAtr
	p.Add(plotter.NewGrid())

	sc, err := plotter.NewScatter(scatterData)
	if err != nil {
		log.Panic(err)
	}

	// Specify style and color for individual points.
	sc.GlyphStyleFunc = func(i int) draw.GlyphStyle {
		_, _, z := scatterData.XYZ(i)
		d := (z - minZ) / (maxZ - minZ)
		rng := maxZ - minZ
		k := d*rng + minZ
		c, err := colors.At(k)
		if err != nil {
			log.Panic(err)
		}
		return draw.GlyphStyle{Color: c, Radius: vg.Points(3), Shape: draw.CircleGlyph{}}
	}

	a, b, c, d := sc.DataRange()

	// fmt.Println(a, b, c, d)

	p.X.Min = a
	p.X.Max = b
	p.Y.Min = c
	p.Y.Max = d

	p.Add(sc)

	// //Create a legend
	// thumbs := plotter.PaletteThumbnailers(colors.Palette(len(scatterData)))
	// for i := len(thumbs) - 1; i >= 0; i-- {
	// 	t := thumbs[i]
	// 	if i != 0 && i != len(thumbs)-1 {
	// 		p.Legend.Add("", t)
	// 		continue
	// 	}
	// 	var val int
	// 	switch i {
	// 	case 0:
	// 		val = int(minZ)
	// 	case len(thumbs) - 1:
	// 		val = int(maxZ)
	// 	}
	// 	p.Legend.Add(fmt.Sprintf("%d", val), t)
	// }

	// This is the width of the legend, experimentally determined.
	const legendWidth = vg.Centimeter

	// Slide the legend over so it doesn't overlap the ScatterPlot.
	p.Legend.XOffs = legendWidth

	img := vgimg.New(520, 520)
	dc := draw.New(img)
	dc = draw.Crop(dc, 0, -legendWidth, 0, 0) // Make space for the legend.
	p.Draw(dc)

	w, err := os.Create(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer w.Close()
	png := vgimg.PngCanvas{Canvas: img}
	if _, err = png.WriteTo(w); err != nil {
		log.Panic(err)
	}
	if err = w.Close(); err != nil {
		log.Panic(err)
	}
}
