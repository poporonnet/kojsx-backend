import { ContestsRepository } from "../../repository/contestRepository.js";
import { Contest } from "../contest.js";

export class ContestService {
  private readonly repository: ContestsRepository;
  constructor(repository: ContestsRepository) {
    this.repository = repository;
  }
  async Exists(d: Contest): Promise<boolean> {
    const all = await this.repository.findAll();
    if (all.isFailure()) {
      return true;
    }

    const iv = all.value.map((v) => {
      return [v.id, v.title];
    });

    return iv
      .map((v) => {
        return v[0] === d.id || v[1] === d.title;
      })
      .includes(true);
  }
}
