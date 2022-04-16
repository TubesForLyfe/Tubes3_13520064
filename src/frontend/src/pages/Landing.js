import React from 'react'
import { Link } from 'react-router-dom'

const Landing = () => {
  return (
    <div>
      Welcome to DNA Application
      <Link to="/add-disease">Add Disease</Link>
      <Link to="/disease-prediction">Disease Prediction</Link>
      <Link to="/detail-prediction">Detail Prediction</Link>
    </div>
  )
}

export default Landing
