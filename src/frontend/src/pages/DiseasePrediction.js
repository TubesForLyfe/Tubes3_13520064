import React, { useState } from 'react'
import './DiseasePrediction.css'
import Axios from 'axios'



const DiseasePrediction = () => {
  const [inputName, setInputName] = useState('');
  const [inputFile, setInputFile] = useState('');
  const [inputPenyakit, setInputPenyakit] = useState('');

  const [hasilTes, setHasilTes] = useState([]);

  const getDiseasePrediction = (e) => {
    e.preventDefault();
    Axios.post(`${process.env.REACT_APP_DNA_API}/get-diseaseprediction`, {
      inputName: inputName,
      inputFile: inputFile,
      inputPenyakit: inputPenyakit,
      tanggal : new Date().toLocaleString().split(" ")[0],
    }).then((response) => {
      setHasilTes(response.data)
    })
  }

  return (
    <div>
      <div id="tesDNA">
        <div>
          <h3> Tes DNA </h3>
          <div className = "column side">
            <p>Nama Penguna:</p>
            <input className="inputBox" type='text' placeholder='Nama pengguna...' 
            onChange={(e) => {
              setInputName(e.target.value)
            }}></input>
          </div>
          <div className = "column middle">
            <p>Sequence DNA:</p>
            <input className="inputBox" type="file" accept=".txt"
            onChange={(e) => {
              setInputFile(e.target.files[0].name)
            }}></input>
          </div>
          <div className = "column side">
            <p>Prediksi Penyakit:</p>
            <input className="inputBox" type='text' placeholder='Nama Penyakit...'
            onChange={(e) => {
              setInputPenyakit(e.target.value)
            }}></input>
          </div>
          <div>
            <button onClick={getDiseasePrediction}>Submit</button>
          </div>
        </div>
        <div id="hasilTes">
          <h3> Hasil </h3>
          {hasilTes.map((val, key) => {
            return (
              <div>
              {val.Status !== -1 &&
                 <div>
                <p className="Same"> {val.TanggalPrediksi} -</p>
                <p className="Same">&nbsp; {val.NamaPasien} -</p>
                <p className="Same">&nbsp; {val.TingkatKemiripan}% -</p>
                <p className="Same">&nbsp; {val.PenyakitPrediksi} -</p>
                {val.Status == 1 && <p className="Same">&nbsp; True</p>}
                {val.Status == 0 && <p className="Same">&nbsp; False</p>}
                </div>
              } 
              {val.Status === -1 &&
                <p className="Same">DNA Tidak Valid!</p>
              }
              </div>
            )
          })}
        </div>
      </div>
      
    </div>
  )
}

export default DiseasePrediction
