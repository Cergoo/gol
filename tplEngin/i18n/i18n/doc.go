// (c) 2013 Cergoo
// under terms of ISC license

/*
Package i18n implementation 
Feature:
    - Load from .json format language resource store.
    - Support tag: include context vars and plural.
    - Support pluggable module as a user functions libraris (a example 'humanize' mod implementation)     
Example: 
    Field {{0}} must be filled {{1}} {{plural appel 1}}
    Good afternoon, Mr.(Mrs.) {{0}}, you have {{f humanByteLong 1}}.
    
See tplEngin/i18n/exaple for more details. 
*/
package i18n
