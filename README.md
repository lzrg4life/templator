# Templator

Someone else has probably already made a better version of this. I made this myself just for fun and for practice.

This takes a json file and a template file as input and uses Go's
templating system to replace values in the template file

```
templater <input json> <template file> <output file>
```

## Example

Given a source.json with:
```
{"Foo":"bar"}
```

And a base.template with:
```
The value of Foo is "{{.Foo}}".
```

Running:
```bash
templator source.json base.template output.txt
```

Produces an output.txt with:
```
The value of Foo is "bar".
```