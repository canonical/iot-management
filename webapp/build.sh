#!/bin/sh

# Build the project
npm run build

# Create the static directory
rm -rf ../static/css
rm -rf ../static/js
cp -R build/static/css ../static/
cp -R build/static/js ../static/
cp build/index.html ../static/app.html

# cleanup
rm -rf ./build
