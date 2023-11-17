import PetRepository from "./petRepository";
import ScheduleRepository from "./scheduleRepository";

export interface Repositories {
  pet: PetRepository;
  schedule: ScheduleRepository;
}

function getRepositories(): Repositories {
  const pet = new PetRepository();
  const schedule = new ScheduleRepository();

  const repositories: Repositories = {
    pet,
    schedule,
  };
  return repositories;
}

export default getRepositories();
