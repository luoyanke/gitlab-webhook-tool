package internal

func PushFeishuCardTmpl() string {
	return `
		{
		  "config": {
			"wide_screen_mode": true
		  },
		  "elements": [
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**项目：**\n{{.projectName}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**推送用户：**\n{{.userName}}"
				  }
				}
			  ]
			},
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**Ref：**\n{{.ref}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**项目地址：**\n{{.webUrl}}"
				  }
				}
			  ]
			},
			{
			  "tag": "hr"
			},
			{
			  "tag": "note",
			  "elements": [{
				"content": "{{.commit}}",
				"tag": "plain_text"
			  }]
			}
		  ],
		  "header": {
			"template": "blue",
			"title": {
			  "content": "{{.title}}",
			  "tag": "plain_text"
			}
		  }
		}
		`
}

func MergeRequestFeishuCardTmpl() string {
	return `{
		  "config": {
			"wide_screen_mode": true
		  },
		  "elements": [
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**项目：**\n{{.projectName}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**合并请求用户：**\n{{.userName}}"
				  }
				}
			  ]
			},
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**source branch：**\n{{.sourceBranch}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**target branch：**\n{{.targetBranch}}"
				  }
				}
			  ]
			},
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**项目地址：**\n{{.webUrl}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": " "
				  }
				}
			  ]
			}
		  ],
		  "header": {
			"template": "blue",
			"title": {
			  "content": "{{.title}}",
			  "tag": "plain_text"
			}
		  }
		}`
}
