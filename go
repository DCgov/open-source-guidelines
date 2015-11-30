#!/bin/bash

function serve { bundle exec jekyll serve --watch --baseurl ''; }

$@
