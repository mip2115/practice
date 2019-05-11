import React, { Component} from "react";

class App extends Component{

  	constructor() {
		super();
		this.state = {
			query: '',
		}
		this.handleChange = this.handleChange.bind(this);
	}

  	handleChange(event) {
  		this.setState({
  			query: event.target.value
  		});
  	}
  	
  	render() {
    	return(
	      <div className="App">
	        <h1> Hello, World??????</h1>
	        <h1> Hello, World!!!</h1>
	        <h1> Hello, World!!!!</h1>
	        <form action="/api/postform" method="POST">
	        <input type="text" name="query" onChange={this.handleChange} />
	        <input type="submit" value="Submit" />
	        </form>
	        
	      </div>
	   	);
	}
  
}

export default App;
