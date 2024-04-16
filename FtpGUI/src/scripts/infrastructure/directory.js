export class Directory {
    constructor(name) {
        this.id = generateUUID();
        this.name = name;
        this.path = ""
        this.parent = null;
        this.directories = [];
        this.files = [];
        this.display = true;
    }
}

export class File {
    constructor(name) {
        this.id = generateUUID();
        this.name = name;
        this.directory = null;
    }

    path() {
        return `${this.directory.path}${this.name}`
    }
}

export class DirectoryTree {
    constructor() {
        this.root = new Directory("root");
        this.root.path = "/"
    }

    findDirectory(directoryId) {
        const result = this.#findDirectory(this.root, directoryId);
        if (result == null)
            throw new Error("Directory id was not found");

        return result;
    }

    insertFile(directoryId, fileName) {
        if (!this.#insertFile(this.root, directoryId, fileName))
            throw new Error("Directory id was not found");
    }

    insertDirectory(parentId, directoryName) {
        if (!this.#insertDirectory(this.root, parentId, directoryName, "/"))
            throw new Error("Directory id was not found");
    }

    removeFile(fileId) {
        if (!this.#removeFile(this.root, fileId))
            throw new Error("File id was not found");
    }

    removeDirectory(directoryId) {
        if (!this.#removeDirectory(this.root, directoryId))
            throw new Error("Directory id was not found")
    }

    toHtml() {
        return `<ul>${this.#toHtml(this.root)}</ul>`;
    }

    #findDirectory(directory, directoryId) {
        if (directory.id == directoryId) {
            return directory;
        }
        else {
            for (const dir of directory.directories) {
                const result = this.#findDirectory(dir, directoryId)
                if (result != null)
                    return result;
            }

            return null;
        }
    }

    #insertFile(directory, directoryId, fileName) {
        if (directory.id == directoryId) {
            const file = new File(fileName);
            file.directory = directory;
            directory.files.push(file);
            return true;
        }
        else {
            for (const dir of directory.directories) {
                if (this.#insertFile(dir, directoryId, fileName))
                    return true;
            }

            return false;
        }
    }

    #insertDirectory(parent, parentId, directoryName, path) {
        if (parent.id == parentId) {
            const directory = new Directory(directoryName);
            directory.path = `${path}${directoryName}/`
            directory.parent = parent;
            parent.directories.push(directory);
            return true;
        }
        else {
            for (const dir of parent.directories) {
                if (this.#insertDirectory(dir, parentId, directoryName, `${path}${dir.name}/`))
                    return true;
            }

            return false;
        }
    }

    #removeFile(directory, fileId) {
        const file = directory.files.find(f => f.id == fileId);

        if (file != undefined) {
            // Remove file from directory
            directory.files = directory.files.filter(f => f.id != fileId)
            return true;
        }
        else {
            for (const dir of directory.directories) {
                if (this.#removeFile(dir, fileId))
                    return true;
            }

            return false;
        }
    }

    #removeDirectory(parent, directoryId) {
        const directory = parent.directories.find(dir => dir.id == directoryId);

        if (directory != undefined) {
            // Remove directory from parent
            parent.directories = parent.directories.filter(dir => dir.id != directoryId)
            return true;
        }
        else {
            for (const dir of parent.directories) {
                if (this.#removeDirectory(dir, directoryId))
                    return true;
            }

            return false;
        }
    }

    #toHtml(directory) {
        let directories = "";
        let files = ""

        for (const dir of directory.directories) {
            directories += this.#toHtml(dir);
        }

        if (directory.display)
            directory.files.forEach(f => {
                files += `<li>${f.name}</li>`;
            });

        return `
            <li>${directory.name}</li>
            <ul>
                ${files}
                ${directories}
            </ul>
        `
    }
}

function generateUUID() {
    // Create a typed array to hold the random values
    var array = new Uint8Array(16);
    // Fill the array with random values
    crypto.getRandomValues(array);

    // Set the version and variant bits
    array[6] = (array[6] & 0x0f) | 0x40; // Version 4
    array[8] = (array[8] & 0x3f) | 0x80; // Variant 10

    // Convert the array to a hexadecimal string
    var hex = Array.from(array).map(b => b.toString(16).padStart(2, '0')).join('');

    // Format the string as a UUID
    return [
        hex.slice(0, 8),
        hex.slice(8, 12),
        hex.slice(12, 16),
        hex.slice(16, 20),
        hex.slice(20, 32)
    ].join('-');
}