// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.650
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Portfolio Lorenz Hohermuth</title><script src=\"https://cdn.jsdelivr.net/npm/@unocss/runtime\"></script><script src=\"https://unpkg.com/htmx.org@1.9.11\" integrity=\"sha384-0gxUXCCR8yv9FM2b+U3FDbsKthCI66oH5IA9fHppQq9DDMHuMauqq1ZHBpJxQ0J0\" crossorigin=\"anonymous\"></script></head><style>\n\t\t@font-face {\n\t\t\tfont-family: UbuntuMono;\n\t\t\tsrc: url(static/UbuntuMono/UbuntuMonoNerdFont-Bold.ttf)\n\t\t}\n\t</style><body class=\"m-0 p-0 font-black\" style=\"font-family: UbuntuMono\"><nav class=\"flex w-full pos-sticky top-0 z-30\"><div class=\"bg-orange-500 w-35 flex justify-center\"><img src=\"/static/robot.png\" height=\"70\"></div><div class=\"container bg-[#3331ee]\"><ul class=\"flex items-center justify-center h-full gap-x-4 p-0 list-none m-0\"><li><a class=\"no-underline text-[#ff5cdb]\" href=\"#home\">Home</a></li><li><a class=\"no-underline text-[#ff5cdb]\" href=\"#projects\">Projects</a></li><li><a class=\"no-underline text-[#ff5cdb]\" href=\"#work\">Work</a></li></ul></div></nav><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main><footer class=\"flex text-xs font-thin\"><div class=\"bg-gray-500 text-gray-400 p-2 w-1/2 flex justify-center items-center\"><p>&copy; 2024 Lorenz Hohermuth All rights reserved.</p></div><div class=\"bg-gray-800 w-1/2 text-neutral-300 flex justify-center items-center\"><p>-- <a class=\" text-neutral-300\" href=\"https://github.com/lorenzhohermuth\">GitHUB</a> -- </p></div></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
