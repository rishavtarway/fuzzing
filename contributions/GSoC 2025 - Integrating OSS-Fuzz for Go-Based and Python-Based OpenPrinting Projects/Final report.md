# Integrating Go- and Python-based OpenPrinting Projects in OSS-Fuzz


* **Year**: 2025
* **Contributor**: Mohammed Imaduddin
* **Organization**: OpenPrinting, The Linux Foundation
* **Mentors**: Till Kamppeter, Jiongchi Yu, George Andrei Iosif
* **Useful Links**:

  * [Project Page](https://summerofcode.withgoogle.com/programs/2025/projects/rBWHXaha)
  * [Source Code for Fuzz Harnesses](https://github.com/OpenPrinting/fuzzing)
  * [OSS-Fuzz](https://github.com/google/oss-fuzz) Projects

    * [goipp](https://github.com/google/oss-fuzz/tree/master/projects/goipp)
    * [ipp-usb](https://github.com/google/oss-fuzz/tree/master/projects/ipp-usb)
    * [pyppd](https://github.com/google/oss-fuzz/tree/master/projects/pyppd)
    * [pycups](https://github.com/google/oss-fuzz/tree/master/projects/pycups)
  * [IPP-over-USB Emulator](https://github.com/OpenPrinting/go-mfp/tree/master/proto/usbip)


---

## Project Context and Significance

OpenPrinting develops and maintains printing-related software, standards, and resources for Linux and other Unix-like systems. It maintains the printer/driver database and key components such as cups-filters and “Printer Applications” which enable IPP-based and driverless printing when used together with CUPS. While previous GSoC work successfully integrated C/C++-based OpenPrinting projects into OSS-Fuzz (achieving 21 bug fixes from 41 discovered issues), **no Go or Python projects had ever been integrated into OSS-Fuzz before this project**.

This project addresses a significant gap in security testing for key printing infrastructure components:

* **goipp** – A Go library maintained by OpenPrinting that implements the Internet Printing Protocol (IPP) for encoding and decoding IPP messages in Go.

* **ipp-usb** – An OpenPrinting userspace daemon that exposes USB printers as local IPP network printers (IPP-over-USB proxy), enabling driverless printing for modern USB devices.

* **pyppd** – A Python library/utility from OpenPrinting for reading, compressing, and decompressing PostScript Printer Description (PPD) files.

* **pycups** – Python bindings to the CUPS API (libcups), allowing Python applications to manage printers and print jobs through CUPS.

The challenge was particularly complex because these projects involve language bindings, hardware-dependent code, and network protocols that are difficult to test without specialized infrastructure.

---

## Previous Work

Building on the foundation established during GSoC 2024, the C/C++ integration work by Jiongchi Yu had already:
- Integrated CUPS and libcups into OSS-Fuzz, discovering 41 security issues with 21 fixes
- Established the [OpenPrinting/fuzzing](https://github.com/OpenPrinting/fuzzing) repository structure
- Created workflows for maintaining fuzzing harnesses externally while integrating with OSS-Fuzz
- Demonstrated the value of continuous fuzzing for critical printing infrastructure

---

## Work Completed During This Project

This project represents the **first successful integration of Go and Python projects into OSS-Fuzz**.

### Technical Achievements by Project

**goipp - Go IPP Implementation**

* Created **5 comprehensive fuzz harnesses** targeting core IPP message processing:

  * `fuzz_decode_bytes.go`: Tests IPP message deserialization with malformed input
  * `fuzz_decode_bytes_ex.go`: Extended deserialization testing with additional parameters
  * `fuzz_collections.go`: Tests complex IPP Collection attribute handling
  * `fuzz_round_trip.go`: Validates encode/decode consistency to catch data corruption
  * `fuzz_tag_extension.go`: Tests custom IPP tag extensions parsing
* **Pull Requests Created**:

  * [OpenPrinting/fuzzing PR #10](https://github.com/OpenPrinting/fuzzing/pull/10)
  * [OpenPrinting/fuzzing PR #11](https://github.com/OpenPrinting/fuzzing/pull/11)
  * [OpenPrinting/fuzzing PR #12](https://github.com/OpenPrinting/fuzzing/pull/12)
  * [OpenPrinting/fuzzing PR #13](https://github.com/OpenPrinting/fuzzing/pull/13)
  * [OpenPrinting/fuzzing PR #18](https://github.com/OpenPrinting/fuzzing/pull/18)
  * [OSS-Fuzz PRs for goipp](https://github.com/google/oss-fuzz/tree/master/projects/goipp)

**ipp-usb - IPP-over-USB Daemon**

* Developed **3 specialized fuzzers** for hardware-dependent code:

  * `daemon_fuzzer.go`: Tests the main daemon with malformed HTTP/IPP requests
  * `usb_fuzzer.go`: Tests USB protocol parsing using mock USB/IP infrastructure
  * `http_client_fuzzer.go`: Tests HTTP client resilience against malicious server responses
* **Major Innovation**: Built a virtual IPP-over-USB emulator using USB/IP protocol, enabling fuzzing of hardware-dependent code paths without requiring physical USB printers
* **Pull Requests Created**:

  * [OpenPrinting/fuzzing PR #26](https://github.com/OpenPrinting/fuzzing/pull/26)
  * [OSS-Fuzz PRs for ipp-usb](https://github.com/google/oss-fuzz/tree/master/projects/ipp-usb)

**pyppd - Python PPD Archive Manager**

* Implemented **7 Python fuzz harnesses** covering file processing and compression:

  * `fuzz_archive.py`: Tests PPD file archiving with malformed input
  * `fuzz_compress.py`: Tests compression/decompression of PPD files
  * `fuzz_compressor.py`: Tests round-trip compression consistency
  * `fuzz_find_files.py`: Tests file discovery with malicious patterns
  * `fuzz_ppd.py`: Tests PPD file parsing with corrupted data
  * `fuzz_read_file_in_syspath.py`: Tests system path file access
  * `fuzz_runner.py`: Tests command-line argument parsing
* **Pull Requests Created**:

  * [OpenPrinting/fuzzing PR #31](https://github.com/OpenPrinting/fuzzing/pull/31)
  * [OpenPrinting/fuzzing PR #32](https://github.com/OpenPrinting/fuzzing/pull/32)
  * [OpenPrinting/fuzzing PR #36](https://github.com/OpenPrinting/fuzzing/pull/36)
  * [OpenPrinting/fuzzing PR #37](https://github.com/OpenPrinting/fuzzing/pull/37)
  * [OSS-Fuzz PRs for pyppd](https://github.com/google/oss-fuzz/tree/master/projects/pyppd)

**pycups - Python CUPS Bindings**

* Created **7 fuzz harnesses** for critical CUPS operations:

  * `fuzz_auth_callback.py`: Tests authentication callback handling
  * `fuzz_buffer_handling.py`: Tests buffer management and memory safety
  * `fuzz_file_handling.py`: Tests file I/O operations
  * `fuzz_ipp_io.py`: Tests IPP request/response processing
  * `fuzz_print_job.py`: Tests job submission with malformed parameters
  * `fuzz_printer_management.py`: Tests printer configuration operations
  * `fuzz_UTF8.py`: Tests string encoding/decoding across all functions
* **Pull Requests Created**:

  * [OpenPrinting/fuzzing PR #27](https://github.com/OpenPrinting/fuzzing/pull/27)
  * [OpenPrinting/fuzzing PR #28](https://github.com/OpenPrinting/fuzzing/pull/28)
  * [OpenPrinting/fuzzing PR #29](https://github.com/OpenPrinting/fuzzing/pull/29)
  * [OpenPrinting/fuzzing PR #35](https://github.com/OpenPrinting/fuzzing/pull/35)
  * [OSS-Fuzz PRs for pycups](https://github.com/google/oss-fuzz/tree/master/projects/pycups)

---

## Impact and Technical Challenges Overcome

### Why This Work Was Difficult

1. Integrated four very different OpenPrinting projects (Go libraries, Python bindings, and utilities) into OSS-Fuzz, each with its own language, build system and fuzzing engine.
2. Dealt with hardware-dependent paths in ipp-usb and pycups by creating a virtual IPP-over-USB printer/emulator to fuzz code that normally requires real devices.
3. Designed seed corpora and mutation strategies tailored to both IPP messages and legacy PPD files to exercise diverse input formats.
4. Maintained all fuzzers externally and wired them into OSS-Fuzz with custom Dockerfiles and build scripts to keep upstream trees clean.

### Quantified Impact

- **Code Coverage Achieved**: [To be measured after September PRs get merged ;(]
- **Projects Newly Covered**: 4 major OpenPrinting projects

### Recommendations for Future Contributors

1. Reuse and extend the existing fuzz harnesses and seed corpora instead of starting from scratch.
2. Continue improving code coverage in hardware-dependent areas of ipp-usb by enhancing the virtual device emulator.
3. Keep fuzzers in an external module and use clear build scripts/Dockerfiles to simplify OSS-Fuzz integration.
4. Triage and upstream any OSS-Fuzz findings promptly to maintainers.
5. Add unit tests alongside fuzzers to ensure deterministic coverage of critical functions.

---

## Future Development

1. Extend fuzzing coverage across all four projects by adding new harnesses and expanding the seed corpora to cover edge cases and rarely used code paths.
2. Enhance the virtual IPP-over-USB emulator to simulate more device behaviors (if possible) for deeper fuzzing of hardware-dependent code.
3. Integrate continuous fuzzing feedback by monitoring OSS-Fuzz reports and filing/triaging issues promptly with maintainers.
4. Broaden test coverage in projects like pycups and pyppd by adding unit tests to complement fuzzing.
5. Document the fuzzing setup so future contributors can easily reproduce, extend, and maintain it.

---

## Acknowledgment

I would like to thank my mentors for their guidance and support throughout this project. **Till Kamppeter** provided invaluable domain expertise on printing protocols and the OpenPrinting architecture. **George-Andrei Iosif** shared his deep knowledge of fuzzing methodologies and the **OSS-Fuzz** infrastructure. **Jiongchi Yu** deserves special recognition for his mentorship, building on his previous GSoC work and offering detailed technical guidance at every step, especially in fuzzing and OSS-Fuzz integration. I am also grateful to **Alexander Pevzner**, the author of **goipp** and **ipp-usb**, for helping me create fuzzing harnesses for these projects and for assisting in the development of the IPP-over-USB emulator using USB/IP. I would also like to thank **Google’s OSS-Fuzz maintainers** for promptly reviewing and merging pull requests. Finally, I thank the wider **OpenPrinting community** and the **Linux Foundation** for maintaining the projects and infrastructure that made this work possible.
