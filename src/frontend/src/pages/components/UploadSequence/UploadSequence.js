import React, { useState, useEffect } from "react"
import "./UploadSequence.css"



let fileValid = false
const UploadSequence = () => {
    const [labelName, setlabelName] = useState("Tidak ada file dipilih")// label below upload button


    const getExtension = (fileName) => {
        return fileName.split('.').pop();
    }

    const setHandler = (event) => {
        var fileName = event.target.files[0].name
        if (fileName != null) {
            if (getExtension(fileName) === "txt") {
                fileValid = true
                setlabelName(fileName)
            }
            else {
                fileValid = false
                setlabelName("Ekstensi file tidak valid")
            }
        }
        else {
            fileValid = false
            setlabelName("Tidak ada file dipilih")
        }
        
    }


    return (

        <div class = "UploadSequence colMargin">
            Upload Sequence:
            <form id="inputPenyakitForm">
                <input id="filePenyakit" onChange={ setHandler } type="file" accept=".txt"/>
                <br/>
                <label id="fileChosen">
                { labelName }
                </label>
            </form>
        </div>

    )

    


}
export default UploadSequence
export const isFileValid = fileValid
