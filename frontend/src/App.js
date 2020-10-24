import React, { useState } from 'react';
import { makeStyles } from "@material-ui/core/styles";
import Button from '@material-ui/core/Button';
import Paper from "@material-ui/core/Paper";
import TextField from "@material-ui/core/TextField";
import IconButton from "@material-ui/core/IconButton";
import SubmitIcon from "@material-ui/icons/Send";
import Chip from "@material-ui/core/Chip";
import './App.css';

const useStyles = makeStyles((theme) => ({
	root: {
		display: 'flex',
		justifyContent: 'center',
		flexWrap: 'wrap',
		listStyle: 'none',
		padding: theme.spacing(0.5),
		margin: 0,
	},
	chip: {
		margin: theme.spacing(0.5),
	},
	iconButton: {
		padding: 10,
		marginTop: 10,
	},
	makeOrderButton: {
		marginTop: 10,
	}
}));

const App = () => {
	const classes = useStyles();

	const [ availablePacks, setAvailablePacks ] = React.useState([ 250, 500, 1000, 2000, 5000 ]);
	const [ newPack, setNewPack ] = React.useState(0);
	const [ quantity, setQuantity ] = useState(0);
	const [ packsNeeded, setPacksNeeded ] = useState([]);
	const [ error, setError ] = useState('');

	const handleDelete = (index) => {
		const reflection = [...availablePacks];
		reflection.splice(index, 1);

		setAvailablePacks(reflection);
	};

	const handleAddPack = () => {
		if (!availablePacks.includes(newPack) && newPack > 0) {
			setAvailablePacks([ ...availablePacks, newPack ])
		} else if (newPack === 0) {
			setError('Pack cannot be 0!');
		} else if (availablePacks.includes(newPack)) {
			setError('Pack already exists!');
		} else {
			setError('Something just went wrong!');
		}
	}

	const makeOrder = async () => {
		setError('');

		const result = await fetch('https://4qyor869oj.execute-api.eu-west-2.amazonaws.com/demo/get-needed-packs', {
			body: JSON.stringify({ packs: availablePacks, quantity }),
			method: 'POST',
			headers: {
				"Content-Type": "application/json"
			}
		});

		if (result.status !== 200) {
			const errorMessage = await result.text();
			setError(errorMessage)
		} else {
			const data = await result.json();
			setPacksNeeded(data.Packs);
		}
	}

	return (
		<div className="App">
			<div className="packs">
				<Paper component="form" className={classes.root} noValidate autoComplete="off">
					<TextField onChange={(e) => setNewPack(+e.target.value)} type="number" id="standard-basic" label="New Pack"/>
					<IconButton onClick={handleAddPack} color="primary" className={classes.iconButton} aria-label="directions">
						<SubmitIcon/>
					</IconButton>
				</Paper>

				<Paper component="ul" className={classes.root}>
					{availablePacks.map((pack, index) => {
						return (
							<li key={index}>
								<Chip
									label={pack}
									onDelete={ () => handleDelete(index)}
									onMouseDown={(e) => {
										e.stopPropagation()
									}}
									className={classes.chip}
								/>
							</li>
						);
					})}
				</Paper>
			</div>
			<br/>
			<TextField onChange={(e) => setQuantity(+e.target.value)} type="number" id="standard-basic" label="Quantity"/>
			<br/>
			{/*<input placeholder="Enter quantity to order" type="number" onChange={(e) => setQuantity(+e.target.value)}/>*/}
			<Button className={classes.makeOrderButton} variant="contained" onClick={makeOrder} color="primary">Make Order</Button>
			<br/>
			{
				packsNeeded.length > 0 && error.length < 1 ? (
					<div>Packs needed: {packsNeeded.toString()}</div>
				) : error.length > 0 ? <div>{error}</div> : ''
			}
		</div>
	);
}

export default App;
