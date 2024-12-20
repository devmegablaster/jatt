// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func DiscordStatus() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"discord-status\" class=\"bg-blue-500 h-2 w-2 rounded-full hover:w-20 space-x-2\"><div id=\"discord-status-text\" class=\"text-transparent\">Loading</div><img src=\"https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fpnghq.com%2Fwp-content%2Fuploads%2Fwhite-discord-logo-png-png-no-watermark-download-57636-1024x745.png&amp;f=1&amp;nofb=1&amp;ipt=56823f049a2f0b78ced821afe592850482f10e8588d164de00e35cd640e7846b&amp;ipo=images\" class=\"h-5 mr-2\"></div><script>\n      function getDiscordStatus() {\n        fetch(\"https://api.lanyard.rest/v1/users/id\")\n          .then(response => response.json())\n          .then(data => {\n            if (data.data.discord_status === \"online\") {\n              document.getElementById(\"discord-status\").classList.remove(\"bg-blue-500\");\n              document.getElementById(\"discord-status\").classList.add(\"bg-green-500\");\n              document.getElementById(\"discord-status-text\").innerText = \"Online\";\n            } else if (data.data.discord_status === \"dnd\") {\n              document.getElementById(\"discord-status\").classList.remove(\"bg-blue-500\");\n              document.getElementById(\"discord-status\").classList.add(\"bg-red-500\");\n              document.getElementById(\"discord-status-text\").innerText = \"Do Not Disturb\";\n            } else if (data.data.discord_status === \"idle\") {\n              document.getElementById(\"discord-status\").classList.remove(\"bg-blue-500\");\n              document.getElementById(\"discord-status\").classList.add(\"bg-yellow-500\");\n              document.getElementById(\"discord-status-text\").innerText = \"Idle\";\n            } else {\n              document.getElementById(\"discord-status\").classList.remove(\"bg-blue-500\");\n              document.getElementById(\"discord-status\").classList.add(\"bg-gray-500\");\n              document.getElementById(\"discord-status-text\").innerText = \"Offline\";\n            }\n          });\n      }\n\n      getDiscordStatus()\n\n      const statusElement = document.getElementById(\"discord-status\");\n      const textElement = document.getElementById(\"discord-status-text\");\n\n      statusElement.addEventListener(\"mouseenter\", () => {\n        const textWidth = textElement.scrollWidth; // Get the full width of the text\n        statusElement.style.width = `${textWidth + 50}px`; // Set it as the new width\n      });\n\n      statusElement.addEventListener(\"mouseleave\", () => {\n        statusElement.style.width = \"1rem\"; // Reset to the initial width\n      });\n  </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
