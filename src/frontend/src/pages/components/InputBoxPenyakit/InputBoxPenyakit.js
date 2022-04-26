import React, { useState, useEffect } from "react"
import "./InputBoxPenyakit.css"

const InputBoxPenyakit = (props) => {
    const setHandler = () => {
        var name = document.getElementById("fieldInputPenyakit").value
        props.func(name != '')
    }

    return (

        <div class = "InputBoxPenyakit colMargin">
            Nama penyakit:
            <form id="InputBoxPenyakitForm">
                <input id="fieldInputPenyakit" type="text" placeholder="Nama Penyakit" onChange={() => setHandler()}/>
            </form>
        </div>

    )

    


}
export default InputBoxPenyakit
