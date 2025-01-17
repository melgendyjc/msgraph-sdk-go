package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type ApplicationGuardBlockFileTransferType int

const (
    // Not Configured
    NOTCONFIGURED_APPLICATIONGUARDBLOCKFILETRANSFERTYPE ApplicationGuardBlockFileTransferType = iota
    // Block clipboard to transfer Image and Text file
    BLOCKIMAGEANDTEXTFILE_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
    // Block clipboard to transfer Image file
    BLOCKIMAGEFILE_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
    // Neither of text file or image file is blocked from transferring
    BLOCKNONE_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
    // Block clipboard to transfer Text file
    BLOCKTEXTFILE_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
)

func (i ApplicationGuardBlockFileTransferType) String() string {
    return []string{"notConfigured", "blockImageAndTextFile", "blockImageFile", "blockNone", "blockTextFile"}[i]
}
func ParseApplicationGuardBlockFileTransferType(v string) (interface{}, error) {
    result := NOTCONFIGURED_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
        case "blockImageAndTextFile":
            result = BLOCKIMAGEANDTEXTFILE_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
        case "blockImageFile":
            result = BLOCKIMAGEFILE_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
        case "blockNone":
            result = BLOCKNONE_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
        case "blockTextFile":
            result = BLOCKTEXTFILE_APPLICATIONGUARDBLOCKFILETRANSFERTYPE
        default:
            return 0, errors.New("Unknown ApplicationGuardBlockFileTransferType value: " + v)
    }
    return &result, nil
}
func SerializeApplicationGuardBlockFileTransferType(values []ApplicationGuardBlockFileTransferType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
