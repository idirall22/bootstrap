# I18n Test harness
- This shows how to use the i18n GoogleSheet tool.
- The GoogleSheets are here:
https://drive.google.com/drive/folders/1SvB8gOFuvD1IF8baU63Wcp0XTOVT1iSV?usp=sharing
see the config sheet for what config to use.
- hugo test website: https://github.com/bep/docuapi
- Same as https://github.com/linode/docs, https://www.linode.com/docs/

# Default layout of a googlesheet file.
- The first row in a googlesheet should be for __languages__.
- The first column in a googlesheet should be for __filename__ or __keys__.

# How to use GoogleSheet tool?
1. Add a new object to config/hugoconfig.yml and populate these fields:

Key|Value
---|---
ID| `1po7GuEo4H04HEPTuG6DFreenBDvJ2QKLUQiIRPKzEBI`
URL|`https://docs.google.com/spreadsheets/d/...`
CSV|`https://docs.google.com/spreadsheets/d/e/.../pub?output=csv`
MERGE|you can choose how to merge googlesheet cells:`cell`, `column` or `row`. 
OUT_DIR| Out folder path i.e: `outputs/content/` and use: `outputs/content/XXX/` and this generate multi language folders i.e: `outputs/content/en/`, `outputs/content/fr/` ...
FILE_NAME| if __empty__ gsheet use default one on googlesheet. use `languages.XXX` to generate file with multi languages name i.e: `languages.en.toml`, `languages.fr.toml`...
EXTENSION| for now only `.toml` is supported. 

2. Install GoogleSheet tool.

3. Run googlesheet -option=hugo

# When should I use `Cell`, `column`, or `row` in __MERGE__ field?
- __cell__: When each cell in googlesheet represent a file.
- __column__: when each column in googlesheet represent a file.
- __row__: when each row in googlesheet represent a file.

# When should I use `/XXX/` tag in __OUT_DIR__?
The __XXX__ tag is used to generate multilanguage __folders__ i.e:
`OUT_DIR`: `"content/xxx/"`.
this will generate folders for each language in content folder i.e:
- content/
    ------------/en/
    ------------/fr/
    ------------/de/
# When should I use `.XXX` tag in __FILE_NAME__?
The __XXX__ tag is used to generate multilanguage __files__ i.e:
`FILE_NAME`: `"i18n/languages.xxx"`.
this will generate a file for each language in i18n folder i.e:
- i18n/
    --------languages.en.toml
    --------languages.fr.toml
    --------languages.de.toml

## Hosting

## Mage & CI

Mage is used for local and CI builds and deploys.

Just tell Github actions to install golang and our bootstrap and then its easy.

