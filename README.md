# galene-translation

Galène is a good WebRTC system, unfortunatelly the UI speak only englisch.

In order have the content in other language we may use translation systems which kann be offered by various KI-Provider. The results are not as good as expected and some textes will not been translated correctly.

The approach for translate-galene is similar but we use allready manually translated textes and replace the english textes as soon the are put to the web page. 

In order to do this we have to use an allready provided translation, for the first time only german. Other language can be provided through providing files with the translated textes.

The translated textes are located within a file translation-XX.js which will be loaded dynamically. A template file (template.txt) which is to be copied to the file translation-XX.js and filled with the translated textes (XX is the 2 charachters for the language, (pt-BR or pt-PT will be pt).

An other file must also be provided galene-XX.css. This file is used in order to correct the size of the Media-activation element.

## installation

Copy the folder translation under the static directory for your galene project.

Insert the line ```<script src="/translation/translate.js" defer></script>``` within the head section of the files 404.html, change-password.html, galene.html and index.html.

After your language is provided you must insert:

```<script src="/translation/translate.js" defer></script>```

at the end of the ```<head>``` part of the files 404.html, change-password.html, galene.html and index.html.

You have also to create a folder translation within the galene static folder and
copy all files to this folder.

## Translating

See also the file translation/tools/README

The file template.txt contain all textes I have extracted from the html files provided by galene and probably most of the textes put to the script galene.js to the html files. Some textes may be not present, an can be inserted within the translation file.



Most of the textes are fix and there are put to the translateList objects list.

```
const translateList = {
"Galène":      "",
" Enable":     "",
...
};
```
The empty string at the rigth side shall contain your translation. If the english string begin or end with a space, respect it within your translated text.

Some text contain fixed and variable parts. These textes will be matched by a regular expression and the variable parts are noted |1|, |2|,... Those Textes are placed within the array of array :

```
const translatePositional = [
["You have been muted by |1|", ""],
"Kick out |1|",                ""],
...
];
```


Your translated Text may use an other order for the variables, and fixed parts for example:

```
var translatePositional = [
["You have been muted by |1|", "Sie wurden von |1| stumm geschaltet"],
"Kick out |1|",                "|1| auswerfen"],
...
["\\(|1| clients\\)",          "(|1| Klienten)"],
];
```

for the last translation the english text contain '(' and ')', there are escaped with '\\'. This due to the need of the regular expression wich look if the text send by galene.js match to this.

## How our scripts works

The file translate.js look for the language configured within the browser and try to load the file translation/galene-XX.css and then the file translation/translation-XX.js. If the javascript file is present a function within translate.js will be called by the translation-XX.js script and the translation will be performed.

The translate.js file will also try to include a CSS file *galene-XX.css*. This file allows you to set the width of the ' Enable' and ' Disable' button to the terms utilized in your language.