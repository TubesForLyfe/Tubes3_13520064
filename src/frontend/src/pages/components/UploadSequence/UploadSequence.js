import React, { useState } from "react"
import "./UploadSequence.css"




const UploadSequence = (props) => {
    let fileValid = false
    var submitButton = document.getElementById("submitButton")

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
                console.log(submitButton.disabled)
            }
        }
        else {
            fileValid = false
            setlabelName("Tidak ada file dipilih")
            console.log(submitButton.disabled)
        }
        console.log(fileValid)
        props.func(fileValid)
        
    }


    return (

        <div class = "UploadSequence colMargin">
            Upload Sequence:
            <form id="uploadSequenceForm">
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
