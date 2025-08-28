# Integrating Go- and Python-based OpenPrinting Projects in OSS-Fuzz

* Year: 2025
* Contributor: Mohammed Imaduddin

**Organization**: OpenPrinting, The Linux Foundation

**Mentors**: Till Kamppeter, Jiongchi Yu, George Andrei Iosif

**Useful Links**:

* [Project Page](need to add link)
* [Source Code for Fuzz Harnesses](https://github.com/OpenPrinting/fuzzing)
* [OSS-Fuzz](https://github.com/google/oss-fuzz) Projects

  * [goipp](https://github.com/google/oss-fuzz/tree/master/projects/goipp)
  * [ipp-usb](https://github.com/google/oss-fuzz/tree/master/projects/ipp-usb)
  * [pyppd](https://github.com/google/oss-fuzz/tree/master/projects/pyppd)
  * [pycups](https://github.com/google/oss-fuzz/tree/master/projects/pycups)
* [IPP-over-USB Emulator](https://github.com/OpenPrinting/go-mfp/tree/master/proto/usbip)

---

## Project Details

This project extends OpenPrinting’s fuzz testing efforts by targeting components written in **Go** and **Python**, complementing the prior work on C/C++ projects. The focus was on **goipp**, **ipp-usb**, **pyppd**, and **pycups**, all of which are widely used in driverless printing and CUPS-based workflows.

The goal was to write fuzzing harnesses, integrate them into the OSS-Fuzz infrastructure, and address the unique challenges of testing language bindings and hardware-dependent code.

---

## Achievements

1. **goipp**: Developed fuzz harnesses for `DecodeBytes`, `DecodeBytesEx`, Collection handling, round-trip Encode/Decode, and TagExtension parsing.

2. **ipp-usb**: Created fuzzers for the daemon, USB protocol parsing, and HTTP client behavior.

3. Built a virtual **IPP-over-USB emulator** using USB/IP, simulating a USB printer to enable fuzzing of hardware-dependent paths and improve coverage without needing physical devices.

4. **pyppd**: Added fuzzers for compression, decompression, file handling, argument parsing, and PPD parsing.

5. **pycups**: Wrote fuzz harnesses for authentication callbacks, buffer handling, IPP I/O, job submission, printer management, and UTF-8 conversions.

6. **Integration**: Continued maintaining fuzzing harnesses in an external repository ([OpenPrinting/fuzzing](https://github.com/OpenPrinting/fuzzing)) with corresponding OSS-Fuzz configs. Also, ensured reproducible local builds and tested workflows using `infra/helper.py`.

---

## Future Development

1. Extend fuzzing seeds for all the fuzzers over all the projects.
2. Improve coverage in `ipp-usb` by expanding emulator capabilities.
3. Expand coverage for other projects, such as goipp, pyppd, and pycups, by targeting low-risk or under-tested areas to increase resilience and catch subtle bugs.
4. Continue triaging OSS-Fuzz reports and working with maintainers for fixes.


---

## Acknowledgment

I would like to extend my deepest gratitude to my mentors throughout this project. **Till Kamppeter** provided invaluable guidance and support throughout, helping me stay on track and prioritize work effectively. **George-Andrei Iosif** contributed his expertise in fuzzing and OSS-Fuzz, which was critical in designing robust fuzzing harnesses and ensuring seamless integration. **Jiongchi Yu** deserves special thanks for his constant encouragement, detailed guidance, and insightful suggestions at every step.

I am also grateful to the OSS-Fuzz maintainers for their feedback on integration, and to the OpenPrinting community for the opportunity to contribute to strengthening the reliability and security of the printing stack.
