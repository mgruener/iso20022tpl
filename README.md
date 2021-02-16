# ISO 20022 Templater

Commandline tool to render templates based on [ISO 20022](https://www.iso20022.org/iso-20022-message-definitions) bank statements.

Primary motivation is to get the digital bank statements of my current bank imported into [Firefly](https://www.firefly-iii.org/) using its [CSV importer](https://docs.firefly-iii.org/csv/).

# Filtering documents

ISO20022 documents can be huge and depending on the bank, contain mostly empty fields. The following command can be used to get a "clean" version of a document to check wich fields can be used
for templating:

```
iso20022tpl json <document> | jq 'walk( if type == "object" then with_entries(select(
      (.value != null) and
      (.value != "") and
      (.value != {}) and
      (.value != []) and
      (.value != "0001-01-01T00:00:00Z") and
      (.value != "0001-01-01Z") and
      (.value != "0001-01Z") and
      (.value != "NOTPROVIDED")
      (.value != 0) and
   )) else . end)'
```

The different ```(.value != ...)``` statements might have to be changed depending on what values can be considered "empty" for the given document.
