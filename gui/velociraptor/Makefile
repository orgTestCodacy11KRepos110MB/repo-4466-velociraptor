all:
	echo "module fake_nodejs_module // Exclude this directory from Go tools\n\ngo 1.16" > 'node_modules/go.mod'
	npm run start

build: FORCE
	echo "module fake_nodejs_module // Exclude this directory from Go tools\n\ngo 1.16" > 'node_modules/go.mod'
	npm run build
	mkdir -p build/static/
	echo > build/static/.keep
	echo > build/.keep

FORCE: ;
