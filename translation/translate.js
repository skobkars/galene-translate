'use strict';

const translate = {
    mkMatchStr : function (s) {
        let reg ='^';
        let a = s;
        for ( let i = 0; i < a.length; i++) {
            if(a[i]==='') {
                continue;
            }
            if ( ! /^\d+$/.test(a[i]) ) {
                reg += a[i];
            } else {
                reg += '(.+?)';
                let idx = a[i];
            }
        }
        reg += '$';
        return reg;
    },
    replace : function (ins, m) {
        let r = '';
        for (let i = 0; i < ins.length; i++) {
            if(ins[i]==='') {
                continue;
            }
            if (  ! /^\d+$/.test(ins[i]) ) {
                r += ins[i];
            } else {
                 r += m[ins[i]];
            }
        }
        return r;
    },
    searchReg : function (input) {
        let strings = translatePositional;
        for (let i = 0; i < strings.length; i++ ) {
            let ins = strings[i][0].split('|');
            let reg = translate.mkMatchStr(ins);
            let m = null;
            if (m = input.match(reg)) {
                let replaced = translate.replace(strings[i][1].split('|'), m);
                return replaced;
            }
        }
        return null;
    },
    chatSystemMessage : function() {
        let box = document.querySelectorAll('#box .message-row');
        if ( !box || box.length == 0 )
            return;
        let last = box[box.length-1];
        let messageContent = last.querySelector('.message-system .message-content div');
        let newContent = "";
        if (messageContent) {
            let content = [''];
            content = messageContent.textContent.split('\n');
            if (content.length > 1) {
                // /help output
                for (var i = 0; i < content.length; i++) {
                    if ( content[i] === "" ) continue;
                    let search = content[i];
                    let part = translateList[search];
                    if ( part )
                       newContent = newContent + part + '\n';
                    else
                       newContent = newContent + content[i] + '\n';
                }
            } else {
                // other system messages
                newContent = translate.replaceTextIn(messageContent, false,false);
                // avoid loop, replace only if changed.
                if ( newContent !== ''  &&  messageContent.textContent !== newContent)
                   newContent = messageContent.textContent;
            }
            // avoid loop, replace only if changed.
            if ( newContent !== '' &&  messageContent.textContent !== newContent)
                messageContent.textContent = newContent;
            return;
        }
        let privatContentP = last.querySelector('.message-private p');
        if ( privatContentP ) {
            let content = privatContentP.textContent;
            if ( content ) {
                newContent = translate.replaceTextIn(privatContentP.parentNode, false,false);
                if ( newContent && newContent != '') 
                    privatContentP.textContent = newContent;
                return;
            }
        }
    },
    replaceTextIn : function(node, includeWhitespaceNodes, checkComposed) {
        let textNodes = [], whitespace = /^\s*$/;
        function getTextNodes(node) {
            if (node.nodeType == 3) {
                let found = false;
                if (includeWhitespaceNodes || !whitespace.test(node.nodeValue)) {
                    if ( node.textContent === '✖') {
                        // ignore toastify button!
                        found = true;
                        return;
                    }
                    if ( !found && translateList[node.textContent] ) {
                        node.textContent = translateList[node.textContent];
                        found = true;
                    }
                    if (!found) {
                        let text = translate.searchReg(node.textContent);
                        if ( text && text !== '' ) {
                            node.textContent = text;
                            found = true;
                        }
                    }
                    if ( ! found )
                        console.log('Miss "'+node.textContent+'"');
                }
            } else {
                for (var i = 0, len = node.childNodes.length; i < len; ++i) {
                    getTextNodes(node.childNodes[i]);
                }
            }
        }
       getTextNodes(node);
    },
    replaceTitle : function() {
        let titles = document.querySelectorAll('*[title]');
        for (let i = 0; i < titles.length; i++) {
            let title = titles[i].getAttribute('title');
            let translated = translateList[title];
            if ( translated && translated != '') {
                titles[i].setAttribute('title', translated);
            }
        }
    },
    replaceAriaLabel : function() {
        let labels = document.querySelectorAll('*[aria-label]');
        for (let i = 0; i < labels.length; i++) {
            let label = labels[i].getAttribute('aria-label');
            let translated = translateList[label];
            if ( translated && translated != '') {
                labels[i].setAttribute('aria-label', translated);
            }
        }
    },
    replaceValue : function() {
        let values = document.querySelectorAll('input[type=submit]');
        for (let i = 0; i < values.length; i++) {
            let value = values[i].getAttribute('value');
            let translated = translateList[value];
            if ( translated && translated != '')
                values[i].setAttribute('value', translated);
        }
    },
    replacePopupText : function(records, observer) {
        for (const record of records) {
            for (const addedNode of record.addedNodes) {
                let popup = document.querySelector('.contextualMenu');
                if (popup) {
                    // Check for composed entry 
                    translate.replaceTextIn(popup, false, true);
                }
                let toast = document.querySelector('.toastify');
                if ( toast )
                    translate.replaceTextIn(toast, false, true);
            }
        }
    },
    replaceTextContentTimed : function() {
        let paraId = [ '#errormessage', '#message' ];
        paraId.forEach(i => {
            let sel = document.querySelector(i);
            if ( sel ) {
                let text = sel.textContent;
                let newText = translateList[text];
                if ( newText )
                   sel.textContent = translateList[text];
                sel.classList.remove('hidden');
            }
       });
    },
    replaceTextContent : function() {
        let errormessage = document.querySelector('#errormessage');
        if ( errormessage )
            errormessage.classList.add('hidden');
        let message = document.querySelector('#message');
        if ( message )
            message.classList.add('hidden');
        setTimeout(replaceTextContentTimed, 100);
    },
    setPlaceHolder : function() {
        let textArea = document.querySelector('#input');
        if ( textArea && textArea.placeholder !== '') {
            let text =textArea.placeholder;
            let newText = translateList[text];
            if ( newText && newText !== text )
                textArea.placeholder = newText;
        }
    },
    replaceText : function() {
        let lang = navigator.language.slice(0,2);
        document.querySelector('html').setAttribute('lang',lang);
        let body = document.querySelector('body');
        translate.replaceTextIn(body);
        translate.replaceValue();
        translate.replaceTitle();
        translate.replaceAriaLabel();
        translate.setPlaceHolder();
        const observerOptions = {
            childList: true,
            subtree: false,
        };
        const observer = new MutationObserver(translate.replacePopupText);
        observer.observe(document.body, observerOptions);
        let chatBox = document.querySelector('#box');
        if ( chatBox ) {
            const chatObserver = new MutationObserver(translate.chatSystemMessage);
            chatObserver.observe(chatBox, {childList:true, subtree:false, characterData:true });
        }
        let input = document.querySelector('#input');
        if ( input ) {
           const inputObserver = new MutationObserver(translate.setPlaceHolder);
           inputObserver.observe(input, {attributes: true});
        }
        let sel = document.getElementById('filterselect');
        if ( sel ) {
            setTimeout(translate.replaceTextIn, 2000, sel);
        }
        let groupform = document.querySelector('#groupform');
        if ( groupform ) {
           let errormessage = document.querySelector('#errormessage');
           groupform.addEventListener('submit',translate.replaceTextContent);
        }
        let pwForm =  document.querySelector('#passwordform');
        if ( pwForm ) {
           pwForm.addEventListener('submit',translate.replaceTextContent);
       }
       let tbl = document.querySelector('#public-groups-table');
       if ( tbl ) {
           setTimeout(translate.replaceTextIn,2000, tbl);
       }
    },
    loadLanguage : function() {
        let lang = navigator.language.slice(0,2);
        let css = document.createElement('link');
        css.rel ='stylesheet';
        css.type ='text/css';
        css.href = '/translation/galene-'+lang+'.css';
        document.head.appendChild(css);
        let script = document.createElement('script');
        script.src = '/translation/translation-'+lang+'.js';
        document.head.appendChild(script);
    },
};

document.addEventListener('DOMContentLoaded',translate.loadLanguage);
