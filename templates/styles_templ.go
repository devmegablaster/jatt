// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func GlobalStyles() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\nbody {\n  margin-left: auto;\n  margin-right: auto;\n  margin-top: 0px;\n  margin-bottom: 0px;\n  display: flex;\n  height: 100vh;\n  width: 91.666667%;\n  flex-direction: column;\n  --tw-bg-opacity: 1;\n  background-color: rgb(29 30 32 / var(--tw-bg-opacity));\n  font-family: ui-sans-serif, system-ui, sans-serif, \"Apple Color Emoji\", \"Segoe UI Emoji\", \"Segoe UI Symbol\", \"Noto Color Emoji\";\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n\n@media (min-width: 640px) {\n  body {\n    width: 66.666667%;\n  }\n}\n\na {\n  text-decoration-line: none;\n}\n\nh1 {\n  font-size: 1.875rem;\n  line-height: 2.25rem;\n}\n\n@media (min-width: 640px) {\n  h1 {\n    font-size: 2.25rem;\n    line-height: 2.5rem;\n  }\n}\n\nh2 {\n  font-size: 1.5rem;\n  line-height: 2rem;\n}\n\n@media (min-width: 640px) {\n  h2 {\n    font-size: 1.875rem;\n    line-height: 2.25rem;\n  }\n}\n\nh3 {\n  font-size: 1.25rem;\n  line-height: 1.75rem;\n}\n\n@media (min-width: 640px) {\n  h3 {\n    font-size: 1.5rem;\n    line-height: 2rem;\n  }\n}\n\nh4 {\n  font-size: 1.125rem;\n  line-height: 1.75rem;\n}\n\n@media (min-width: 640px) {\n  h4 {\n    font-size: 1.25rem;\n    line-height: 1.75rem;\n  }\n}\n\nh5 {\n  font-size: 1rem;\n  line-height: 1.5rem;\n}\n\n@media (min-width: 640px) {\n  h5 {\n    font-size: 1.125rem;\n    line-height: 1.75rem;\n  }\n}\n\np {\n  font-size: 0.875rem;\n  line-height: 1.25rem;\n}\n\n@media (min-width: 640px) {\n  p {\n    font-size: 1rem;\n    line-height: 1.5rem;\n  }\n}\n\nol {\n  margin: 0px;\n  list-style-position: outside;\n  list-style-type: decimal;\n  padding-left: 1rem;\n}\n\nol li {\n  margin: 0px;\n  list-style-position: outside;\n  list-style-type: decimal;\n  padding-left: 1rem;\n}\n\nol li:not(:last-child) {\n  margin-bottom: 0.5rem;\n}\n\nul {\n  margin: 0px;\n  list-style-position: outside;\n  list-style-type: disc;\n  padding-left: 1rem;\n}\n\nul li {\n  margin: 0px;\n  list-style-position: outside;\n  list-style-type: disc;\n  padding-left: 1rem;\n}\n\nul li:not(:last-child) {\n  margin-bottom: 0.5rem;\n}\n\nfigure {\n  margin-top: 1rem;\n  margin-bottom: 1rem;\n  margin-left: 0px;\n  margin-right: 0px;\n  display: flex;\n  width: 100%;\n  flex-direction: column;\n  align-items: center;\n  justify-content: center;\n}\n\nfigcaption {\n  margin-top: 0.5rem;\n  font-size: 0.875rem;\n  line-height: 1.25rem;\n  --tw-text-opacity: 1;\n  color: rgb(107 114 128 / var(--tw-text-opacity));\n}\n\nfigure img {\n  margin-left: auto;\n  margin-right: auto;\n  max-width: 75%;\n}\n\nhr {\n  margin-top: 1rem;\n  margin-bottom: 1rem;\n  width: 100%;\n  border-top-width: 1px;\n  --tw-border-opacity: 1;\n  border-color: rgb(107 114 128 / var(--tw-border-opacity));\n}\n\n.active-nav {\n  text-decoration-line: underline;\n  text-underline-offset: 2px;\n}\n\n.hljs {\n  background-color: #2E2E33 !important;\n  display: block;\n  overflow-x: auto;\n  border-radius: 0.375rem;\n  padding: 0.5rem;\n}\n\n.top-container {\n  display: flex;\n  flex-direction: column;\n  align-items: flex-start;\n  justify-content: space-between;\n  padding-top: 0.5rem;\n  padding-bottom: 1rem;\n}\n\n@media (min-width: 640px) {\n  .top-container {\n    flex-direction: row;\n    align-items: center;\n  }\n}\n\n.nav-title {\n  font-size: 1.25rem;\n  line-height: 1.75rem;\n  font-weight: 700;\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n\n@media (min-width: 640px) {\n  .nav-title {\n    font-size: 1.5rem;\n    line-height: 2rem;\n  }\n}\n\n.nav-container {\n  display: flex;\n  align-items: center;\n  justify-content: center;\n}\n\n.nav-container > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-x-reverse: 0;\n  margin-right: calc(1rem * var(--tw-space-x-reverse));\n  margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));\n}\n\n.nav-container {\n  padding-top: 0.75rem;\n  padding-bottom: 0.75rem;\n}\n\n@media (min-width: 640px) {\n  .nav-container > :not([hidden]) ~ :not([hidden]) {\n    --tw-space-x-reverse: 0;\n    margin-right: calc(2rem * var(--tw-space-x-reverse));\n    margin-left: calc(2rem * calc(1 - var(--tw-space-x-reverse)));\n  }\n\n  .nav-container {\n    padding-left: 0.25rem;\n    padding-right: 0.25rem;\n  }\n}\n\n.nav-container a {\n  font-size: 1rem;\n  line-height: 1.5rem;\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n  text-underline-offset: 2px;\n}\n\n.nav-container a:hover {\n  text-decoration-line: underline;\n}\n\n.post-matter {\n  margin-left: auto;\n  margin-right: auto;\n  display: flex;\n  width: 91.666667%;\n  flex-direction: column;\n}\n\n.post-matter > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-y-reverse: 0;\n  margin-top: calc(0.5rem * calc(1 - var(--tw-space-y-reverse)));\n  margin-bottom: calc(0.5rem * var(--tw-space-y-reverse));\n}\n\n@media (min-width: 640px) {\n  .post-matter {\n    width: 66.666667%;\n  }\n}\n\n.post-matter h1 {\n  margin: 0px;\n  font-size: 1.5rem;\n  line-height: 2rem;\n}\n\n@media (min-width: 640px) {\n  .post-matter h1 {\n    font-size: 1.875rem;\n    line-height: 2.25rem;\n  }\n}\n\n.post-matter p {\n  margin: 0px;\n}\n\n.post-seperator {\n  margin-top: 25px !important;\n  height: 0.125rem;\n  width: 100%;\n  --tw-bg-opacity: 1;\n  background-color: rgb(107 114 128 / var(--tw-bg-opacity));\n}\n\n.tags-container {\n  display: flex;\n  align-items: center;\n}\n\n.tags-container > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-x-reverse: 0;\n  margin-right: calc(0.5rem * var(--tw-space-x-reverse));\n  margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));\n}\n\n.tags-container-listing {\n  display: flex;\n  align-items: center;\n}\n\n.tags-container-listing > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-x-reverse: 0;\n  margin-right: calc(0.5rem * var(--tw-space-x-reverse));\n  margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));\n}\n\n@media (min-width: 640px) {\n  .tags-container-listing {\n    position: absolute;\n    top: 1rem;\n    right: 1rem;\n  }\n}\n\n.tag {\n  border-radius: 0.375rem;\n  --tw-bg-opacity: 1;\n  background-color: rgb(82 82 82 / var(--tw-bg-opacity));\n  padding-left: 0.5rem;\n  padding-right: 0.5rem;\n  padding-top: 0.25rem;\n  padding-bottom: 0.25rem;\n  font-size: 0.75rem;\n  line-height: 1rem;\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n\n@media (min-width: 640px) {\n  .tag {\n    font-size: 0.875rem;\n    line-height: 1.25rem;\n  }\n}\n\nfooter {\n  width: 100%;\n  padding-top: 2rem;\n  padding-bottom: 2rem;\n  text-align: center;\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n\nfooter a {\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n  text-decoration-line: underline;\n  text-underline-offset: 2px;\n}\n\na:visted {\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n  text-decoration-line: underline;\n  text-underline-offset: 2px;\n}\n\n.content-raw a {\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n  text-decoration-line: underline;\n  text-underline-offset: 2px;\n}\n\n.content-raw a:visted {\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n  text-decoration-line: underline;\n  text-underline-offset: 2px;\n}\n\n.content-raw {\n  margin-left: auto;\n  margin-right: auto;\n  width: 91.666667%;\n  flex-grow: 1;\n}\n\n@media (min-width: 640px) {\n  .content-raw {\n    width: 66.666667%;\n  }\n}\n\n.home-container {\n  margin-top: auto;\n  margin-bottom: auto;\n  display: flex;\n  flex-grow: 1;\n  flex-direction: column;\n  align-items: center;\n  justify-content: center;\n}\n\n.home-container > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-y-reverse: 0;\n  margin-top: calc(1rem * calc(1 - var(--tw-space-y-reverse)));\n  margin-bottom: calc(1rem * var(--tw-space-y-reverse));\n}\n\n.home-container {\n  text-align: center;\n}\n\n.home-image {\n  height: 6rem;\n  width: 6rem;\n  border-radius: 9999px;\n}\n\n@media (min-width: 640px) {\n  .home-image {\n    height: 8rem;\n    width: 8rem;\n  }\n}\n\n.name-description-container {\n  display: flex;\n  flex-direction: column;\n  align-items: center;\n}\n\n.name-description-container > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-y-reverse: 0;\n  margin-top: calc(0.5rem * calc(1 - var(--tw-space-y-reverse)));\n  margin-bottom: calc(0.5rem * var(--tw-space-y-reverse));\n}\n\n.home-name {\n  margin: 0px;\n  font-size: 1.875rem;\n  line-height: 2.25rem;\n  font-weight: 700;\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n\n@media (min-width: 640px) {\n  .home-name {\n    font-size: 2.25rem;\n    line-height: 2.5rem;\n  }\n}\n\n.home-description {\n  margin: 0px;\n  font-size: 1.125rem;\n  line-height: 1.75rem;\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n\n.home-social-container {\n  display: flex;\n}\n\n.home-social-container > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-x-reverse: 0;\n  margin-right: calc(1rem * var(--tw-space-x-reverse));\n  margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));\n}\n\n.home-social-container a {\n  height: 1.25rem;\n  width: 1.25rem;\n  --tw-text-opacity: 1;\n  color: rgb(59 130 246 / var(--tw-text-opacity));\n}\n\n.home-social-container a:hover {\n  --tw-text-opacity: 1;\n  color: rgb(29 78 216 / var(--tw-text-opacity));\n}\n\n@media (min-width: 640px) {\n  .home-social-container a {\n    height: 1.5rem;\n    width: 1.5rem;\n  }\n}\n\n.home-button-container {\n  display: flex;\n}\n\n.home-button-container > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-x-reverse: 0;\n  margin-right: calc(1.5rem * var(--tw-space-x-reverse));\n  margin-left: calc(1.5rem * calc(1 - var(--tw-space-x-reverse)));\n}\n\n.home-button-container {\n  padding-top: 1rem;\n}\n\n.home-button-container a {\n  border-radius: 0.375rem;\n  --tw-bg-opacity: 1;\n  background-color: rgb(65 66 68 / var(--tw-bg-opacity));\n  padding-left: 1rem;\n  padding-right: 1rem;\n  padding-top: 0.5rem;\n  padding-bottom: 0.5rem;\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n\n.listing-container {\n  margin-left: auto;\n  margin-right: auto;\n  display: flex;\n  width: 91.666667%;\n  flex-grow: 1;\n  flex-direction: column;\n}\n\n.listing-container > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-y-reverse: 0;\n  margin-top: calc(1rem * calc(1 - var(--tw-space-y-reverse)));\n  margin-bottom: calc(1rem * var(--tw-space-y-reverse));\n}\n\n@media (min-width: 640px) {\n  .listing-container {\n    width: 66.666667%;\n  }\n}\n\n.listing-item-container {\n  position: relative;\n  display: flex;\n  width: auto;\n  flex-direction: column;\n}\n\n.listing-item-container > :not([hidden]) ~ :not([hidden]) {\n  --tw-space-y-reverse: 0;\n  margin-top: calc(0.25rem * calc(1 - var(--tw-space-y-reverse)));\n  margin-bottom: calc(0.25rem * var(--tw-space-y-reverse));\n}\n\n.listing-item-container {\n  border-radius: 0.375rem;\n  --tw-bg-opacity: 1;\n  background-color: rgb(46 46 51 / var(--tw-bg-opacity));\n  padding-left: 1.5rem;\n  padding-right: 1.5rem;\n  padding-top: 1rem;\n  padding-bottom: 1rem;\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n\n.listing-item-container h3 {\n  margin: 0px;\n  font-size: 1.25rem;\n  line-height: 1.75rem;\n  font-weight: 700;\n}\n\n@media (min-width: 640px) {\n  .listing-item-container h3 {\n    font-size: 1.5rem;\n    line-height: 2rem;\n  }\n}\n\n.listing-item-container p {\n  margin: 0px;\n  font-size: 0.875rem;\n  line-height: 1.25rem;\n  --tw-text-opacity: 1;\n  color: rgb(107 114 128 / var(--tw-text-opacity));\n}\n\n#discord-status {\n  display: flex;\n  height: 1rem;\n  width: 1rem;\n  align-items: center;\n  justify-content: center;\n  overflow: hidden;\n  transition-property: all;\n  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\n  transition-duration: 150ms;\n}\n\n#discord-status:hover {\n  height: 1.75rem;\n}\n\n#discord-status-text {\n  transform: translateY(1rem);\n  width: -moz-fit-content;\n  width: fit-content;\n  white-space: nowrap;\n  padding: 0.5rem;\n}\n\n#discord-status:hover #discord-status-text {\n  transform: translateY(0);\n  --tw-text-opacity: 1;\n  color: rgb(255 255 255 / var(--tw-text-opacity));\n}\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
