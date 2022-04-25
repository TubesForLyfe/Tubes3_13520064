import React, { useState, useEffect } from "react"
import "../SubmitButton/SubmitButton"
import SubmitButton from "../SubmitButton/SubmitButton"
import InputBoxPenyakit from "../InputBoxPenyakit/InputBoxPenyakit"
import "./Container.css"
import UploadSequence from "../UploadSequence/UploadSequence"

const Container = () => {

    return (

        <div class = "ContainerBody ">
            <h3>Input Penyakit</h3>
            <InputBoxPenyakit/>
            <UploadSequence/>
            <br/>
            <SubmitButton/>
        </div>

    )

    


}
export default Container
