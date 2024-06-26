// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.650
package page

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/lorenzhohermuth/portfolio/view/layout"
import "github.com/lorenzhohermuth/portfolio/view/component"

func ShowHome(arr []component.CarouselEntry, index int, events []component.Event) templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"bg-[#cdb4f6] pt-1\"><div class=\"md:flex justify-center\"><div id=\"home\" class=\"text-white py-2 pr-3 pl-6 bg-[#6266ec] my-4 mr-4 rounded-r-4 md:rounded-4 md:m-4 md:w-164\"><h2>About this Website </h2><p>The Purpose of this website is to try a new Stack with GO and HTMX with the goal of using a Minimal amount of Code. And to refresh my old Portfolio to a newer State.</p><p>Normally, my Portfolio is seriously out of date. So I have written this one with the GitHub API and a small parser, which parser the Markdown Files in my Repo into the Timeline and the Carousel. They are the files in the Interactive folder. </p><a href=\"https://github.com/LorenzHohermuth/portfolio\" class=\"text-[#ff5cdb]\">Repo</a><p>Yes, I could have written a GitHub workflow to Automatically make restart the site, but I wanted to learn more about the go Context. (Plus Pipelines are a Pain in the Ass).</p></div></div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = component.Banner("Stack", "", "#6266ec", "#b9f301").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"sm:flex justify-center py-12\"><div class=\"px-4 sm:w-150\"><div class=\"grid grid-cols-3 gap-1 gap-y-2 text-sm md:text-base\"><div class=\"col-span-2 rounded bg-[#ff5cdb] flex justify-center items-center py-3 px-6\">Language</div><div class=\"rounded bg-[#ff5cdb] flex justify-center items-center py-3 px-6\">GO</div><div class=\"text-white rounded bg-[#6266ec] flex justify-center items-center py-3 px-6\">HTMX</div><div class=\"text-white col-span-2 rounded bg-[#6266ec] flex justify-center items-center py-3 px-6\">Frontend Famework</div><div class=\"col-span-2 rounded bg-[#b9f301] flex justify-center items-center py-3 px-6\">CSS Famework</div><div class=\"rounded bg-[#b9f301] flex justify-center items-center py-3 px-6\">Uno CSS</div><div class=\"text-white rounded bg-[#6266ec] flex justify-center items-center py-3 px-6\">Templ</div><div class=\"text-white col-span-2 rounded bg-[#6266ec] flex justify-center items-center py-3 px-6\">HTML Templating</div><div class=\"text-white col-span-2 rounded bg-[#3331ee] flex justify-center items-center py-3 px-6\">HTTP Famework</div><div class=\"text-white rounded bg-[#3331ee] flex justify-center items-center py-3 px-6\">Echo</div></div></div></div></section><section id=\"projects\" class=\"bg-[#ff5cdb] pb-7\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = component.Banner("Projects", "󰣪", "#b9f301", "#ff5cdb").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex justify-center items-center mt-10\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = component.Carousel(arr, index).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></section><section id=\"work\" class=\"bg-[#fee7c5]\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = component.Banner("Work", "", "#1733d2", "#ff5cdb").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"p-3 h-full\"><div class=\"bg-[#1733d2] rounded-8 flex justify-center items-center py-16 sm:p-24\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = component.Timeline(events).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></section>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
