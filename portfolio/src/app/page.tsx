import Image from 'next/image'
import SkillIcon from '../Icons/SkillsIcon'

export default function Home() {
  return (
    <main className='py-10 font-medium'>
      <div className='flex items-center flex-col'>
        <Image src="/Images/BSI-Portrait_Lorenz_Q.jpg" 
        alt='a image of myself'
        width={800}
        height={800}
        className='rounded-full h-32 w-32 shadow-md'>  
        </Image>
        <div className=' text-center text-zinc-500 text-xl my-4'>
          <div className='text-zinc-700'>Lorenz</div>
          <div>Hohermuth</div>
        </div>
      </div>
      <nav className='py-3 flex justify-evenly px-5'>
        <a href="" className=' text-center text-zinc-400 text-sm'>
          <div className=' bg-purple-400 h-16 w-16 rounded-full text-3xl flex justify-center items-center drop-shadow-md shadow-purple-400'>
            ⚡
          </div>
          <div className='mt-1'>
            Skills
          </div>
        </a>
        
        <a href="" className=' text-center text-zinc-400 text-sm'>
          <div className=' bg-green-400 h-16 w-16 rounded-full text-3xl flex justify-center items-center drop-shadow-md '>
            👷
          </div>
          <div className='mt-1'>
            Projects
          </div>
        </a>

        <a href="" className=' text-center text-zinc-400 text-sm'>
          <div className=' bg-orange-400 h-16 w-16 rounded-full text-3xl flex justify-center items-center drop-shadow-md '>
            📧
          </div>
          <div className='mt-1'>
            Contact
          </div>
        </a>

      </nav>

    </main>
  )
}
