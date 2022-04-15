import { BrowserRouter as Router, Switch, Route, Redirect } from "react-router-dom";
import Axios from "axios";

import Landing from "./pages/Landing";
import AddDisease from "./pages/AddDisease";
import DiseasePrediction from "./pages/DiseasePrediction";
import DetailPrediction from "./pages/DetailPrediction";

const RouteManager = () => {
    Axios.defaults.withCredentials = true;

    return (
        <Router>
            <Switch>
                <Route path="/" exact component={Landing} />
                <Route path="/add-disease" component={AddDisease} />
                <Route path="/detail-prediction" component={DetailPrediction} />
                <Route path="/disease-prediction" component={DiseasePrediction} />
                <Route path="*">
                    <Redirect to="/" />
                </Route>
            </Switch>
        </Router>
    )
}

export default RouteManager;