import React, { useState } from 'react';
import './App.css';

function App() {
	const [ packs, setPacks ] = useState([ 250, 500, 1000, 2000, 5000 ]);
	const [ newPack, setNewPack ] = useState([]);
	const [ quantity, setQuantity ] = useState(0);
	const [ packsNeeded, setPacksNeeded ] = useState([]);

	const [ loading, setLoading ] = useState(false);
	const [ error, setError ] = useState('');

	const addPack = () => {
		const diff = newPack.filter(np => !packs.includes(np));
		setPacks([ ...packs, ...diff ])
	}

	const handleOnChangePack = (e) => {
		const data = e.target.value;

		const newpacks = data.split(',').filter(p => p.trim().length && p.trim() > 0).map(pack => +pack);
		setNewPack(newpacks)
	}

	const handleOnChangeQuantity = (e) => setQuantity(+e.target.value)

	const makeOrder = async () => {
		setLoading(true);
		setError('');

		const result = await fetch('https://4qyor869oj.execute-api.eu-west-2.amazonaws.com/demo/get-needed-packs', {
			body: JSON.stringify({ packs, quantity }),
			method: 'POST',
			headers: {
				"Content-Type": "application/json"
			}
		});

		setLoading(false);

		if (result.status !== 200) {
			const errorMessage = await result.text();
			setError(errorMessage)
		} else {
			const data = await result.json();
			console.log(data);
			setPacksNeeded(data.Packs);
		}
	}

	return (
		<div className="App">
			<div className="packs">
				Available Packs:
				<br/>
				{packs.toString()}
			</div>
			<input type="text" onChange={handleOnChangePack}/>
			<button onClick={addPack}>Add packs</button>
			<br/>
			<input placeholder="Enter quantity to order" type="number" onChange={handleOnChangeQuantity}/>
			<button onClick={makeOrder}>Order Quantity</button>
			<br/>
			{
				packsNeeded.length > 0 ? (
					<div>Packs needed: {packsNeeded.toString()}</div>
				) : error.length > 0 ? <div>{error}</div> : ''
			}
		</div>
	);
}

export default App;
