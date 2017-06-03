# DashF

[![Build Status](https://travis-ci.org/cdornsife/dashf.svg?branch=master)](https://travis-ci.org/cdornsife/dashf)

Unmarshall A file, STDIN or a URL into a golang struct.

## What is DashF?

DashF means the flag -f.  When a CLI wishes to accept a resource (filename/directory/url) they typically use the -f flag. DashF provides your cli a consistent way to unmarshall user-defined resources into golang structs. Most CLI tools process the -f flag in the same manner, but the code is typically buried and not easily re-usable.

## Doesn't Viper/envconfig/pflag do this?
No. Viper and envconfig provide a great libraries for configuring your application. Viper and pflag can accept the input but provide no mechanism for processing the resource. 

## [Try the example](example/README.md)