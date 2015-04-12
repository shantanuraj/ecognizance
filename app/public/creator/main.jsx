var ArticlesInput = React.createClass({
	render:function() {
		return (
			<div>
				<div className="form-group">
				  <input type="text" className="form-control" placeholder="Enter title" name={"Articles." + this.props.data + ".Title"} />
				</div>
				<div className="form-group">
				  <input type="text" className="form-control" placeholder="Enter author name" name={"Articles." + this.props.data + ".Author"} />
				</div>
				<div className="form-group">
				  <input type="email" className="form-control" placeholder="Enter author email" name={"Articles." + this.props.data + ".Email"} />
				</div>
				<div className="form-group">
				  <textarea className="form-control" rows="11" placeholder="Enter post" name={"Articles." + this.props.data + ".Post"}></textarea>
				</div>
			</div>
		);
	}
});

var EcohustleInput = React.createClass({
	render: function() {
		return (
			<div>
				<div className="form-group">
	            	<input type="text" className="form-control" placeholder="Enter image URL" name={"Ecohustle." + this.props.data + ".Image"} />
	          	</div>
	          	<div className="form-group">
	            	<input type="text" className="form-control" placeholder="Enter image caption" name={"Ecohustle." + this.props.data + ".Caption"} />
	          	</div>
          	</div>
		);
	}
});

var ArtcilesInputGroup = React.createClass({
	render:function() {
		var rows = [];
		for (var i = 0; i < this.props.data; i++) {
		    rows.push(<ArticlesInput data={i} key={i}/>);
		}
		return <div>{rows}</div>;
	}
})

var EcohustleInputGroup = React.createClass({
	render:function() {
		var rows = [];
		for (var i = 0; i < this.props.data; i++) {
		    rows.push(<EcohustleInput data={i} key={i}/>);
		}
		return <div>{rows}</div>;
	}
})

function updateSupplementaryGroup() {
	var i = document.getElementsByClassName('supplementary-selector')[0].value;
	React.render(
		<ArtcilesInputGroup data={i}/>,
		document.getElementById('supplementary-input')
	);
}

function updateEcohustleGroup() {
	var i = document.getElementsByClassName('ecohustle-selector')[0].value;
	React.render(
		<EcohustleInputGroup data={i}/>,
		document.getElementById('ecohustle-input')
	);
}

