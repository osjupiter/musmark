# musmark

**musmark** is a simple templating tool that manages templates, data, and their output in a single file. It can be used for various purposes such as commands, SQL queries, and code generation, making it highly effective for streamlining everyday repetitive tasks.


**musmark** is a Go CLI tool that manages templates and data in a single Markdown-compatible file, rendering Mustache templates with YAML data.



## ✨ Features

- 🗂️ Manage templates and data in a single file
- 🧠 Support for Mustache template engine
- 📄 Flexible data structure with YAML format
- ⚡ easy to use single binary with Go




## 📦 Installation

```bash
go install github.com/osjupiter/musmark@latest
```

Or for development, build locally:

```bash
go build -o musmark
```



## 🧪 Usage

### Example Input File (`collect_logs.md`)

The input file combines a template and data in Markdown format:

````markdown
template
===
```mustache
#!/bin/bash
{{#servers}}
rsync -az {{user}}@{{host}}:{{log_path}} /logs/{{name}}/
{{/servers}}
tar czf /logs/all_{{env}}_{{date}}.tar.gz /logs/*/
```

data
===
```yaml
env: production
date: "20250407"
servers:
  - name: app1
    host: app1.prod.example.com
    user: syslog
  - name: app2
    host: app2.prod.example.com
    user: syslog
  - name: db
    host: db.prod.example.com
    user: dbadmin
```
````

This template will generate commands like:

````markdown
result
===
```
#!/bin/bash
rsync -az syslog@app1.prod.example.com: /logs/app1/
rsync -az syslog@app2.prod.example.com: /logs/app2/
rsync -az dbadmin@db.prod.example.com: /logs/db/

tar czf /logs/all_production_20250407.tar.gz /logs/*/
```
````

The tool intelligently handles existing result blocks:
- If no result block exists: Adds one at the end
- If result block exists: Updates only that section while preserving the rest

### Execute Command

```bash
# Generate and show complete template with result
musmark collect_logs.md

# Generate and execute the commands directly
musmark -r collect_logs.md | bash

# Preview the generated commands
musmark -r collect_logs.md
```



## 🛠 Future Ideas

- [ ] JSON support
- [ ] Go template support
- [ ] Multiple template sections in a file
- [ ] Enhanced error handling

---

## 📄 License

MIT License
