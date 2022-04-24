import React, { useState } from 'react'
import Axios from 'axios'

const DetailPrediction = () => {
  const [input, setInput] = useState('');
  const [index, setIndex] = useState(1);
  const [predictionList, setPredictionList] = useState([]);

  const getDetailPrediction = (e) => {
    e.preventDefault();
    Axios.post(`${process.env.REACT_APP_DNA_API}/get-detailprediction`, {
      input: input
    }).then((response) => {
      setPredictionList(response.data)
    })
  }

  return (
    <div id='tesDNA'>
      <h3> Detail Prediksi DNA </h3>
      <input type='text' placeholder='Input tanggal atau nama penyakit'
        onChange={(e) => {
          setInput(e.target.value)
        }}
      />
      <button onClick={getDetailPrediction}>Process</button>
      <h3>Hasil</h3>
      <div>
        {predictionList.map((val, key) => {
          return (
            <div className='flex-row ContainerBody'>
              <p>{index}.</p>
              <p className='margin-left-4'>{val.TanggalPrediksi} -</p>
              <p className='margin-left-4'>{val.NamaPasien} -</p>
              <p className='margin-left-4'>{val.PenyakitPrediksi} -</p>
              <p className='margin-left-4'>{val.TingkatKemiripan}% -</p>
              {val.Status == 1 && <p className='margin-left-4'>True</p>}
              {val.Status == 0 && <p className='margin-left-4'>False</p>}
            </div>
          )
        })}
      </div>
    </div>
  )
}

export default DetailPrediction
