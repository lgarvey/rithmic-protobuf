#!/bin/bash

PROTO_DIR="../proto"
GO_PACKAGE="github.com/lgarvey/dom/proto"

# Find all .proto files in the directory
find "$PROTO_DIR" -name "*.proto" | while read -r file; do
    # Check if go_package already exists
    if grep -q "option go_package" "$file"; then
        # Replace existing go_package with the correct one
        sed -i "s|option go_package = .*;|option go_package = \"$GO_PACKAGE\";|" "$file"
        echo "Updated go_package in: $file"
    else
        # If `syntax = "proto2";` is missing, add it
        if ! grep -q 'syntax = "proto3";' "$file"; then
            echo 'syntax = "proto3";' | cat - "$file" > temp && mv temp "$file"
            echo "Added missing syntax declaration to: $file"
        fi

        # Ensure go_package is added right after the syntax line
        awk -v gp="option go_package = \"$GO_PACKAGE\";" '
            /syntax = "proto3";/ { print; print gp; next }
            { print }
        ' "$file" > temp && mv temp "$file"

        echo "Added go_package to: $file"
    fi
done

echo "✅ All .proto files updated!"
