import React from 'react';
import './Card.css';

const Card = ({pokemon}) => {
  return (
    <div className="card">
        <div className="cardImg">
            <img src={pokemon.sprites.other["official-artwork"].front_default} alt="ポケモン画像" className="pokemonImg"/>
        </div>
        <h3>{pokemon.jaName}</h3>
        <div className="cardTypes">
            <div>タイプ</div>
            <span className="typeName">
                {pokemon.type}
            </span>
        </div>
        <div className="cardInfo">
            <div className="cardData">
                <div className="title">重さ：{pokemon.weight}</div>
            </div>
            <div className="cardData">
                <div className="title">高さ：{pokemon.height}</div>
            </div>
            <div className="cardData">
                <div className="title">能力：{pokemon.ability}</div>
            </div>
        </div>
    </div>
  )
}

export default Card