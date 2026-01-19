package gsdfaux

import (
	"fmt"
	"os"

	"github.com/soypat/gsdf/glbuild"
)

func renderIRMF(cfg RenderConfig, s glbuild.Shader3D) error {
	irmfWatch := stopwatch()
	bb := s.Bounds()
	min, max := bb.Min, bb.Max
	header := glbuild.IRMFHeader{
		IRMF:      "1.0",
		Materials: []string{"material0"},
		Min:       [3]float32{min.X, min.Y, min.Z},
		Max:       [3]float32{max.X, max.Y, max.Z},
		Units:     "mm", // Default to mm as per spec
	}
	var objects []glbuild.ShaderObject
	_, objects, err := glbuild.NewDefaultProgrammer().WriteIRMF(cfg.IRMFOutput, s, header)
	if err != nil {
		return fmt.Errorf("writing IRMF: %s", err)
	}
	for i := range objects {
		if objects[i].IsBindable() {
			return fmt.Errorf("bindable objects unsupported for IRMF shader generation: %v", objects[i])
		}
	}
	filename := "IRMF"
	if fp, ok := cfg.IRMFOutput.(*os.File); ok {
		filename = fp.Name()
	}
	fmt.Printf("[%s] wrote %s\n", irmfWatch().String(), filename)
	return nil
}
