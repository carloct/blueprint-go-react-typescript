import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import Home from "@/components/pages/Home";
import About from "@/components/pages/About";
import "./App.css";

function App() {
  return (
    <Router>
      <div className="App">
        <nav>
          <Link to="/">Home</Link>
          <Link to="/about">About</Link>
        </nav>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
