package constraints

import "mime/multipart"
import "fmt"
import "strings"

func NewFileConstraint(options map[string]interface{}) * FileConstraint {
    opt := map[string]interface{}{
        "min-size": 0, // in bytes
        "max-size": int(^uint(0) >> 1), // in bytes
        "allowed-mime-types": []string{},
        "allowed-extensions": []string{},
    }

    for k, v := range options {
        opt[k] = v
    }

    constraint := FileConstraint{}
    constraint.SetOptions(opt)
    constraint.SetErrorMessage("Problem to upload file")

    return &constraint
}

type FileConstraint struct
{
    errorMessage string
    options map[string]interface{}
}

func (self * FileConstraint) SetErrorMessage(errorMessage string) {
    self.errorMessage = errorMessage
}

func (self FileConstraint) GetErrorMessage() string {
    return self.errorMessage
}

func (self * FileConstraint) SetOptions(options map[string]interface{}) {
    self.options = options
}

func (self *FileConstraint) Validate(value interface{}) bool {
    if _, ok := value.(*multipart.FileHeader); ok == false {
        return false
    }

    v := value.(*multipart.FileHeader)

    options := self.options

    size := int(v.Size)
    minSize := options["min-size"].(int)
    if minSize > 0 && size < minSize {
        self.SetErrorMessage(fmt.Sprintf("The minimum size for the file is '%v' bytes", minSize))
        return false
    }

    maxSize := options["max-size"].(int)
    if maxSize > 0 && size > maxSize {
        self.SetErrorMessage(fmt.Sprintf("The maximum size for the file is '%v' bytes", maxSize))
        return false
    }

    fileName := v.Filename
    index := strings.LastIndex(fileName, ".")
    if index == -1 {
        self.SetErrorMessage("File must have extension")
        return false
    }

    allowedExtensions:= options["allowed-extensions"].([]string)
    if len(allowedExtensions) > 0 {
        self.SetErrorMessage(fmt.Sprintf("The allowed extension(s) for the file is/are '%s'", strings.Join(allowedExtensions, ", ")))
        ext := fileName[index+1:]
        inArray := false
        for _, v := range allowedExtensions {
            if v == ext {
                inArray = true
                break
            }
        }

        if inArray == false {
            return false
        }
    }

    allowedMimeTypes := options["allowed-mime-types"].([]string)
    if len(allowedMimeTypes) > 0 {
        self.SetErrorMessage(fmt.Sprintf("The allowed mime type(s) for the file is/are '%s'", strings.Join(allowedMimeTypes, ", ")))

        contentTypes := v.Header["Content-Type"]

        inArray := false
        for _, v := range allowedMimeTypes {
            for _, contentType := range contentTypes {
                if v == contentType {
                    inArray = true
                    break
                }
            }

            if inArray {
                break
            }
        }

        if inArray == false {
            return false
        }
    }

    return true
}
