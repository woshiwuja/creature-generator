const $cursor = document.querySelector('#cursor');

// Listeners
document.body.addEventListener('mousemove', onMouseMove);

// Move cursor
function onMouseMove(e) {
  TweenMax.to($cursor, .1, {
    x: e.pageX - 5,
    y: e.pageY - 7
  })
}

const commands = String.raw`anarchy2036@anarchy2036.xyz/:~ cd go-homepage
anarchy2036@anarchy2036.xyz/go-homepage:~ ls`;

const beep = String.raw`


          _____                    _____                    _____                    _____          
         /\    \                  /\    \                  /\    \                  /\    \         
        /::\    \                /::\    \                /::\    \                /::\    \        
       /::::\    \              /::::\    \              /::::\    \              /::::\    \       
      /::::::\    \            /::::::\    \            /::::::\    \            /::::::\    \      
     /:::/\:::\    \          /:::/\:::\    \          /:::/\:::\    \          /:::/\:::\    \     
    /:::/__\:::\    \        /:::/__\:::\    \        /:::/__\:::\    \        /:::/__\:::\    \    
   /::::\   \:::\    \      /::::\   \:::\    \      /::::\   \:::\    \      /::::\   \:::\    \   
  /::::::\   \:::\    \    /::::::\   \:::\    \    /::::::\   \:::\    \    /::::::\   \:::\    \  
 /:::/\:::\   \:::\ ___\  /:::/\:::\   \:::\    \  /:::/\:::\   \:::\    \  /:::/\:::\   \:::\____\ 
/:::/__\:::\   \:::|    |/:::/__\:::\   \:::\____\/:::/__\:::\   \:::\____\/:::/  \:::\   \:::|    |
\:::\   \:::\  /:::|____|\:::\   \:::\   \::/    /\:::\   \:::\   \::/    /\::/    \:::\  /:::|____|
 \:::\   \:::\/:::/    /  \:::\   \:::\   \/____/  \:::\   \:::\   \/____/  \/_____/\:::\/:::/    / 
  \:::\   \::::::/    /    \:::\   \:::\    \       \:::\   \:::\    \               \::::::/    /  
   \:::\   \::::/    /      \:::\   \:::\____\       \:::\   \:::\____\               \::::/    /   
    \:::\  /:::/    /        \:::\   \::/    /        \:::\   \::/    /                \::/____/    
     \:::\/:::/    /          \:::\   \/____/          \:::\   \/____/                  ~~          
      \::::::/    /            \:::\    \               \:::\    \                                  
       \::::/    /              \:::\____\               \:::\____\                                 
        \::/____/                \::/    /                \::/    /                                 
         ~~                       \/____/                  \/____/                                  


# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #`;

const samurai = String.raw`

█████╗ ███╗   ██╗ █████╗ ██████╗  ██████╗██╗  ██╗██╗   ██╗██████╗  ██████╗ ██████╗  ██████╗ 
██╔══██╗████╗  ██║██╔══██╗██╔══██╗██╔════╝██║  ██║╚██╗ ██╔╝╚════██╗██╔═████╗╚════██╗██╔════╝ 
███████║██╔██╗ ██║███████║██████╔╝██║     ███████║ ╚████╔╝  █████╔╝██║██╔██║ █████╔╝███████╗ 
██╔══██║██║╚██╗██║██╔══██║██╔══██╗██║     ██╔══██║  ╚██╔╝  ██╔═══╝ ████╔╝██║ ╚═══██╗██╔═══██╗
██║  ██║██║ ╚████║██║  ██║██║  ██║╚██████╗██║  ██║   ██║   ███████╗╚██████╔╝██████╔╝╚██████╔╝
╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝ ╚═════╝ ╚═════╝  ╚═════╝`;

let blink = document.querySelector('.blink');
const code = document.querySelector('.code');

const RandomNumber = (min, max) => {
	return Math.floor(Math.random() * max) + min
};

const Delay = (time) => {
	return new Promise((resolve) => setTimeout(resolve, time))
};

const ResetTerminal = () => {
	code.innerHTML = '<span class="blink">█</span>';
	blink = document.querySelector('.blink');
};

const RenderString = characters => {
	blink.insertAdjacentHTML('beforeBegin', characters);
};

const TypeString = async characters => {
	for(const character of characters.split('')) {
		await Delay(RandomNumber(100, 200));
		RenderString(character);
	}
}

const DrawLines = async ( characters, min = 50, max = 500 ) => {
	for(const line of characters.split('\n')) {
		await Delay(RandomNumber(min, max));
		RenderString(`${line}\n`);
	}
}

const DrawCommands = async commands => {
	for( const line of commands.split('\n')){
		// Seperate the directory and the command
		const [currentDir, command] = line.split(':~ ');

		// Print the directory, type the command and finish with new line
		RenderString(`${currentDir}:~ `);
		await TypeString(command);
		RenderString('\n');
	}
}


// Start the code
(async()=> {
	await DrawCommands("/:~ ssh anarchy2036@anarchy2036.xyz -p 22");
	await Delay(1000);
	RenderString("anarchy2036@anarchy2036.xyz password:");
	await Delay(5000);
	RenderString("\n cd go-homepage");
	await DrawCommands(commands);
	RenderString('\n go run server/main.go\n\n');
	await DrawCommands('anarchy2036@anarchy2036.xyz/go-homepage:~ go run server/main.go');
	await DrawLines( beep );
	await TypeString("\n\nServer running.");
	await Delay(3000);
	ResetTerminal();
	await DrawCommands('anarchy2036@anarchy2036.xyz:~ KEY=3db7ca618243da1ba3bc76ab14bcf07b go run server/main.go');
	await DrawLines(samurai);
})();


