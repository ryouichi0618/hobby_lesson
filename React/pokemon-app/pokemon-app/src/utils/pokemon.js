export const getAllPokemon = (url) => {
    return new Promise((resolve, reject) => {
        fetch(url)
            .then(res => res.json())
            .then(data => {
                resolve(data)
            });
    });
};

export const getPokemon = (url) => {
    let resData = new Promise((resolve, reject) => {
        fetch(url)
        .then(res => res.json())
        .then(async data => {
            let jaName = await getPokemonJaName(data);
            data.jaName = jaName;

            let resPokemonTypes = data.types.map((v) => {
                let typesURL = v.type.url;
                return typesURL;
            });

            let resPokemonAbilities = data.abilities.map((v) => {
                let abilitiesURL = v.ability.url;
                return abilitiesURL;
            });
            data.type = await loadPokemonType(resPokemonTypes);
            data.ability = await loadPokemonAbility(resPokemonAbilities);
            data.weight = (Math.floor(data.weight) / 10) + 'kg';
            data.height = (Math.floor(data.height) / 10) + 'm'
            return data;
        })
        .then(data => resolve(data));
    });
    return resData;
};

const getPokemonJaName = (data) => {
    console.log(data)
    return new Promise((resolve, reject) => {
        fetch(data.species.url)
            .then(res => res.json())
            .then(result => {
                return resolve(result.names.find(name => name.language.name === "ja").name);
            })
    });
}

const loadPokemonType = async (data) => {
    let _pokemonType = await Promise.all(
        data.map(async (pokemon) => {
        let pokemonTypeDetail = await getPokemonType(pokemon);
        let jaName = pokemonTypeDetail.names.find(name => name.language.name === "ja").name;
        return jaName;
    })
    );
    return _pokemonType.join(" / ");
};

const getPokemonType = (url) => {
    let resData = new Promise((resolve, reject) => {
        fetch(url)
            .then(res => res.json())
            .then(data => resolve(data));
    });
    return resData;
};

const loadPokemonAbility = async (data) => {
    let _pokemonType = await Promise.all(
        data.map(async (pokemon) => {
        let pokemonTypeDetail = await getPokemonAbility(pokemon);
        let jaName = pokemonTypeDetail.names.find(name => name.language.name === "ja").name;
        return jaName;
    })
    );
    return _pokemonType.join(" / ");
};

const getPokemonAbility = (url) => {
    let resData = new Promise((resolve, reject) => {
        fetch(url)
            .then(res => res.json())
            .then(data => resolve(data));
    });
    return resData;
};