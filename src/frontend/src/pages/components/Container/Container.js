import React, { useState, useEffect } from "react"
import "../SubmitButton/SubmitButton"
import SubmitButton from "../SubmitButton/SubmitButton"
import InputBoxPenyakit from "../InputBoxPenyakit/InputBoxPenyakit"
import "./Container.css"
import UploadSequence from "../UploadSequence/UploadSequence"

const Container = () => {
    const [fileValid, setfileValid] = useState(false)
    const [namaValid, setnamaValid] = useState(false)

    useEffect(() => {
        disable(false, false)
    }, [])
    

    const disable = (fvalid, nvalid) => {
        var button = document.getElementById("submitButton")
        button.disabled = !(fvalid && nvalid)
    }

    const setHandler_nama = (bool) => {
        var elmt = document.getElementById("uploadSequenceForm")
        setnamaValid(bool)
        disable(fileValid, bool)
    }
    const setHandler_file = (bool) => {
        setfileValid(bool)
        disable(bool, namaValid)
    }



    return (

        <div class = "ContainerBody ">
            <h3>Input Penyakit</h3>
            <InputBoxPenyakit func={(arg) => setHandler_nama(arg)}/>
            <UploadSequence func={(arg) => setHandler_file(arg)}/>
            <br/>
            <SubmitButton fileValid={fileValid} namaValid={namaValid} />
            <div id="message"></div>
        </div>

    )

    


}
export default Container
